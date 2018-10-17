package users

import (
	"encoding/json"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Type password exists so we can encrypt while Unmarshaling
type passwordHash []byte

// User type is the model for the user data
type User struct {
	Email     string       `json:"email"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Password  passwordHash `json:"password"`
}

// DbUser is the stub of the user database
var DbUser = map[string]User{}

// UnmarhsalJSON for password lets us unmarshal directly to a hash
func (p *passwordHash) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Panicln(err)
	}

	*p = newPassword
	return nil
}

// AddNewUser adds User to database
func AddNewUser(u User) error {
	if _, userAlreadyInDb := DbUser[u.Email]; userAlreadyInDb {
		return errors.New("User is already in database")
	}
	DbUser[u.Email] = u
	return nil
}
