package auth

import (
	"net/http"
	"webapp_template/backend/sessions"

	uuid "github.com/satori/go.uuid"
)

func createSession(w http.ResponseWriter, email string) {
	newUUID := uuid.Must(uuid.NewV4())
	createCookie(w, newUUID)
	sessions.DbSession[newUUID.String()] = email
}

func createCookie(w http.ResponseWriter, uuid uuid.UUID) {
	cookie := http.Cookie{Name: "SessionID",
		Value: uuid.String()}

	http.SetCookie(w, &cookie)
}

func isSomeoneSignedIn(r *http.Request) bool {
	if _, err := r.Cookie("SessionID"); err == nil {
		return true
	}
	return false
}
