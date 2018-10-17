package auth

import (
	"encoding/json"
	"net/http"
	"webapp_template/backend/users"
)

// Signup signs the user up from a post request
// containing the JSON of the user
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var u users.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = users.AddNewUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
