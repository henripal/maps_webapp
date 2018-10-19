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
	email, err := sessions.GetEmailFromSession(cookie.Value)
	if err != nil {
		http.Error(w, "Email not Found in Session Database", http.StatusBadRequest)
	}
	u, err := users.GetUser(email)
	if err != nil {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(u)
}
