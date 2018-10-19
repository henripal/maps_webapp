package sessions

import (
	"log"
	"testing"
	"webapp_template/backend/users"
)

func init() {
	if err := users.InitializeDB(); err != nil {
		log.Fatalln(err)
	}
}

func Test_AddSession(t *testing.T) {
	ssid := "ha"
	email := "daffy@gmail.com"

	if err := AddNewSession(ssid, email); err != nil {
		t.Fatal(err)
	}

	if email, _ := GetEmailFromSession(ssid); email != "daffy@gmail.com" {
		t.Fatalf("user was not added")
	}

	if err := AddNewSession(ssid, "otheremail"); err == nil {
		t.Fatalf("double session not detected")
	}

	if err := DeleteSession(ssid); err != nil {
		t.Fatalf("unable to delete session")
	}

}
