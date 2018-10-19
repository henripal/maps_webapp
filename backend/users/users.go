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

// SigninUser is the type to be decoded for signin
type SigninUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidUser returns nil if the user is verified and authenticated
func ValidUser(u SigninUser) error {
	userFromDb, err := GetUser(u.Email)
	if err != nil {
		return errors.New("Email not recognized")
	}
	err = bcrypt.CompareHashAndPassword(userFromDb.Password, []byte(u.Password))
	if err != nil {
		return errors.New("Password Incorrect")
	}

	return nil
}

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
	if _, err := GetUser(u.Email); err == nil {
		return errors.New("User is already in database")
	}
	err := addUser(u)
	return err
}
