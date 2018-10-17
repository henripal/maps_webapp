package auth

import (
	"encoding/json"
	"net/http"
	"webapp_template/backend/users"
)

// Signup signs the user up from a post request
// containing the JSON of the user
func Signup(w http.ResponseWriter, r *http.Request) {
	if isSomeoneSignedIn(r) {
		http.Error(w, "Someone is already signed in.", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		newUser := decodeUserAndAddToDb(w, r)
		createSession(w, newUser.Email)
	}
}

func decodeUserAndAddToDb(w http.ResponseWriter, r *http.Request) users.User {
	var u users.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = users.AddNewUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return u
}
