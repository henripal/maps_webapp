package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"webapp_template/backend/sessions"
	"webapp_template/backend/users"

	"golang.org/x/crypto/bcrypt"
)

func Test_signup(t *testing.T) {
	reader := strings.NewReader(`{"email":"daffy@gmail.com","firstName":"daffy","lastName":"duck","password":"quack"}`)
	req, err := http.NewRequest("POST", "http://example.com/foo", reader)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	Signup(res, req)

	act := users.DbUser["daffy@gmail.com"]
	err = bcrypt.CompareHashAndPassword(act.Password, []byte("quack"))
	if err != nil {
		t.Fatalf("passwords don't match")
	}
	if act.FirstName != "daffy" {
		t.Fatalf("first names don't match")
	}
}
func Test_sessionCreation(t *testing.T) {
	email := "daffy@gmail.com"
	res := httptest.NewRecorder()
	createSession(res, email)

	request := &http.Request{Header: http.Header{"Cookie": res.HeaderMap["Set-Cookie"]}}

	cookie, err := request.Cookie("SessionID")
	if err != nil {
		t.Fatal(err)
	}

	uuidString := cookie.Value

	if _, ok := sessions.DbSession[uuidString]; !ok {
		t.Fatal("Sessions not Updated")
	}
}
