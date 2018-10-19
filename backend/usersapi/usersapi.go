package usersapi

import (
	"encoding/json"
	"net/http"
	"webapp_template/backend/sessions"
	"webapp_template/backend/users"
)

// User shows the email associated with a sessionid
func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		http.Error(w, "Not logged in.", http.StatusBadRequest)
		return
	}
	email := sessions.DbSession[cookie.Value]
	u, err := users.GetUser(email)
	if err != nil {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(u)
}
