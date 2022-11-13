package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

var DB *pop.Connection


type User struct {
	ID			uuid.UUID 	`db:"id"`
	FirstName	string		`db:"first_name"`
	LastName	string		`db:"last_name"`
	Password	string		`db:"password"`
	Email		string		`db:"email"`
	Verified	bool		`db:"verified"`
	Admin		bool		`db:"admin"`
}

type Error struct {
	Message		string
}

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
