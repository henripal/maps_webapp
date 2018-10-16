package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

// Type password exists so we can encrypt while Unmarshaling
type passwordHash []byte

type user struct {
	Email     string       `json:"email"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Password  passwordHash `json:"password"`
}

var dbUser = map[string]user{}

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

// backendPortEnvVariable returns the port on which the backend
// will run. By default (dev) ":8888"
func backendPortEnvVariable() string {
	env := os.Getenv("BACKENDPORT")
	if env == "" {
		return "8888"
	}
	return env
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var u user
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = addNewUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func addNewUser(u user) error {
	if _, userAlreadyInDb := dbUser[u.Email]; userAlreadyInDb {
		return errors.New("User is already in database")
	}
	dbUser[u.Email] = u
	return nil
}

func main() {
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":"+backendPortEnvVariable(), cors.Default().Handler(http.DefaultServeMux))
}
