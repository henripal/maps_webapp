package auth

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"webapp_template/backend/sessions"
	"webapp_template/backend/users"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	if err := users.InitializeDBUsers(); err != nil {
		log.Fatalln(err)
	}
}

func Test_signup(t *testing.T) {
	reader := strings.NewReader(`{"email":"daffy@gmail.com","firstName":"daffy","lastName":"duck","password":"quack"}`)
	req, err := http.NewRequest("POST", "http://example.com/foo", reader)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	Signup(res, req)

	act, err := users.GetUser("daffy@gmail.com")
	if err != nil {
		t.Fatalf("Unable to get mock user")
	}
	err = bcrypt.CompareHashAndPassword(act.Password, []byte("quack"))
	if act.FirstName != "daffy" {
		t.Fatalf("first names don't match")
	}

	err = users.DeleteUser("daffy@gmail.com")
	if err != nil {
		t.Fatalf("Error in deleting user")
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
