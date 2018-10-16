package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_signup(t *testing.T) {
	reader := strings.NewReader(`{"email":"daffy@gmail.com","firstName":"daffy","lastName":"duck","password":"quack"}`)
	req, err := http.NewRequest("POST", "http://example.com/foo", reader)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	signup(res, req)

	act := dbUser["daffy@gmail.com"]
	err = bcrypt.CompareHashAndPassword(act.Password, []byte("quack"))
	if err != nil {
		t.Fatalf("passwords don't match")
	}
	if act.FirstName != "daffy" {
		t.Fatalf("first names don't match")
	}
}
