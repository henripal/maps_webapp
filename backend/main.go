package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type user struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

var dbUser = map[string]user{}

// backendPortEnvVariable returns the port on which the backend
// will run. By default (dev) ":8888"
func backendPortEnvVariable() string {
	env := os.Getenv("BACKENDPORT")
	if env == "" {
		return "8888"
	}
	return env
}

func data(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "WHAT IS THE TEMPLATE????")
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
			fmt.Println(err.Error())
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
	http.HandleFunc("/data", data)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":"+backendPortEnvVariable(), cors.Default().Handler(http.DefaultServeMux))
}
