package auth

import (
	"fmt"
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
	if err := users.InitializeDB(); err != nil {
		log.Fatalln(err)
	}
}

func Test_signup(t *testing.T) {
	email := "not@areal.email"
	firstName := "not"
	lastName := "areal"
	password := "quack"

	jsonString := `{"email":"` + email +
		`","firstName":"` + firstName +
		`","lastName":"` + lastName +
		`","password":"` + password + `"}`

	fmt.Println(jsonString)

	reader := strings.NewReader(jsonString)
	req, err := http.NewRequest("POST", "http://example.com/foo", reader)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	Signup(res, req)

	act, err := users.GetUser(email)
	if err != nil {
		t.Fatalf("Unable to get mock user")
	}
	err = bcrypt.CompareHashAndPassword(act.Password, []byte(password))
	if act.FirstName != firstName {
		t.Fatalf("first names don't match")
	}

	err = users.DeleteUser(email)
	if err != nil {
		t.Fatalf("Error in deleting user")
	}

	err = sessions.DeleteSessionFromEmail(email)
	if err != nil {
		t.Fatalf("Error in deleting session")
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

	if _, err := sessions.GetEmailFromSession(uuidString); err != nil {
		t.Fatal("Sessions not Updated")
	}

	if err := sessions.DeleteSession(uuidString); err != nil {
		t.Fatal("unable to delete session")
	}

}
