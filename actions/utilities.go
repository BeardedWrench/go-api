package actions

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"onyx/models"
	"unicode/utf8"

	"github.com/gobuffalo/buffalo"
)

func pkcsPad(input []byte, blockSize int) []byte {
	padLen := blockSize - (len(input) % blockSize)
	padded := make([]byte, len(input)+padLen)
	copy(padded, input)

	padding := padded[len(input):]
	for i, _ := range padding {
		padding[i] = byte(padLen)
	}
	return padded
}

func pkcsUnpad(input []byte, blockSize int) ([]byte, error) {
	if len(input)%blockSize != 0 {
		return nil, fmt.Errorf("input length %d not divisible by block size %d", len(input), blockSize)
	}
	if len(input) < blockSize {
		return nil, fmt.Errorf("input length %d is smaller than block size %d", len(input), blockSize)
	}
	padLen := int(input[len(input)-1])
	if padLen <= 0 || padLen > blockSize {
		return nil, fmt.Errorf("invalid padding length %d", padLen)
	}
	for pos, byte := range input[len(input)-padLen:] {
		if int(byte) != padLen {
			return nil, fmt.Errorf("padding byte %d at pos %d is not the same as padding length %d", byte, pos, padLen)
		}
	}
	return input[:len(input)-padLen], nil
}

func Encrypt(plaintext []byte, secretKey []byte) ([]byte, error) {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize())
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCEncrypter(aes, iv)
	padded := pkcsPad(plaintext, aes.BlockSize())
	// put the IV at the beginning of the ciphertext
	encrypted := make([]byte, len(iv)+len(padded))
	copy(encrypted[:len(iv)], iv)
	cbc.CryptBlocks(encrypted[len(iv):], padded)

	return encrypted, nil
}

// Encrypts a UTF-8 string using secretKey.  The output will be encoded in
// base64 to support storing in the database as a "text" value.
func EncryptString(plaintext string, secretKey []byte) (string, error) {
	ciphertext, err := Encrypt([]byte(plaintext), secretKey)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func Decrypt(ciphertext []byte, secretKey []byte) ([]byte, error) {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	if (len(ciphertext) % aes.BlockSize()) > 0 {
		return nil, fmt.Errorf("input length %d is not a multiple of blocksize %d", len(ciphertext), aes.BlockSize())
	}

	iv := ciphertext[:aes.BlockSize()]
	cbc := cipher.NewCBCDecrypter(aes, iv)
	ciphertext = ciphertext[len(iv):]
	// decrypt in-place
	cbc.CryptBlocks(ciphertext, ciphertext)
	unpadded, err := pkcsUnpad(ciphertext, aes.BlockSize())
	if err != nil {
		return nil, err
	}
	return unpadded, err
}

func DecryptString(ciphertext string, secretKey []byte) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	data, err := Decrypt(decoded, secretKey)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(data) {
		return "", errors.New("decrypted string is not valid UTF-8 data")
	}
	return string(data), nil
}

func OutputError(c buffalo.Context, errorCode int, errorMessage error) error {
	return c.Render(errorCode, r.JSON(models.Error{Message: errorMessage.Error()}))
}

