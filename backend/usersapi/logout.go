package usersapi

import (
	"net/http"
)

// Logout deletes the cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		http.Error(w, "Not logged in.", http.StatusBadRequest)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
