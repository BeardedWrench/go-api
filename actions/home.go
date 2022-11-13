package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"onyx/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)
type UserResource struct{}

var key = []byte{43, 230, 20, 255, 145, 150, 115, 232, 100, 236, 124, 113, 94, 68, 224, 55, 242, 198, 181, 141, 184, 141, 60, 0, 42, 175, 12, 173, 210, 142, 104, 6}

func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome!"}))
}

func (ur UserResource) List(c buffalo.Context) error {
	users := []models.User{}
	error := models.DB.All(&users)

	if error != nil {
		return OutputError(c, http.StatusBadRequest, error)
	}

	return c.Render(http.StatusOK, r.JSON(users))
}

func (ur UserResource) FindUser(c buffalo.Context) error {
	user := models.User{}
	err := models.DB.Find(&user, c.Param("id"))

	if err != nil {
		return OutputError(c, http.StatusBadRequest, err)
	}
	decryptedPassword, decryptedErr := DecryptString(user.Password, key)

	if decryptedErr != nil {
		return OutputError(c, http.StatusBadRequest, decryptedErr)
	}

	//TODO: should we decrypt the password when retrieving a user?
	user.Password = decryptedPassword

	return c.Render(http.StatusOK, r.JSON(user))
}

func (ur UserResource) Create(c buffalo.Context) error {

	body, error := ioutil.ReadAll(c.Request().Body)

	if error != nil{
		fmt.Printf("Error reading body: %v", error)
		return error
	}

	user := &models.User{}
	json.Unmarshal([]byte(body), user)

	encryptedPassword, encryptedErr := EncryptString(user.Password, key)

	if encryptedErr != nil {
		return OutputError(c, http.StatusBadRequest, encryptedErr)
	}

	user.Password = encryptedPassword
	newUser, err := models.DB.ValidateAndCreate(user, "ID")

	if err != nil {
		return OutputError(c, http.StatusBadRequest, err)
	}
	models.DB.ValidateAndSave(newUser, "ID")

	return c.Render(http.StatusCreated, r.JSON(user))
}

func (ur UserResource) Update(c buffalo.Context) error {
	id := c.Param("id")
	body, error := ioutil.ReadAll(c.Request().Body)

	if error != nil{
		fmt.Printf("Error reading body: %v", error)
		return error
	}

	user := &models.User{ID: uuid.FromStringOrNil(id)}
	json.Unmarshal([]byte(body), user)

	encryptedPassword, encryptedErr := EncryptString(user.Password, key)

	if encryptedErr != nil {
		return OutputError(c, http.StatusBadRequest, encryptedErr)
	}

	user.Password = encryptedPassword
	newUser, err := models.DB.ValidateAndUpdate(user, "ID")

	if err != nil {
		return OutputError(c, http.StatusBadRequest, err)
	}
	models.DB.ValidateAndSave(newUser, "ID")

	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "User with id: " + id + " updated."}))
}

func (ur UserResource) Delete(c buffalo.Context) error {
	id := c.Param("id")

	user := &models.User{ID: uuid.FromStringOrNil(id)}

	err := models.DB.Destroy(user)
	if err != nil {
		return OutputError(c, http.StatusBadRequest, err)
	}

	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "User with id: " + id + " deleted."}))
}