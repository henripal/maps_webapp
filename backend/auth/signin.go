package auth

import (
	"encoding/json"
	"net/http"
	"webapp_template/backend/users"
)

// Signin signs the user in from a post request
// containing the JSON of the user
func Signin(w http.ResponseWriter, r *http.Request) {
	if isSomeoneSignedIn(r) {
		http.Error(w, "Someone is already signed in.", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		newUser := decodeSigninUser(w, r)
		loginErr := users.ValidUser(newUser)
		if loginErr != nil {
			http.Error(w, loginErr.Error(), http.StatusBadRequest)
		}
		createSession(w, newUser.Email)
	}
}

func decodeSigninUser(w http.ResponseWriter, r *http.Request) users.SigninUser {
	var u users.SigninUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return u
}
