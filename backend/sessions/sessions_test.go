package sessions

import (
	"testing"
)

func Test_AddSession(t *testing.T) {
	ssid := "ha"
	email := "daffy@gmail.com"

	if err := AddNewSession(ssid, email); err != nil {
		t.Fatal(err)
	}

	if DbSession[ssid] != "daffy@gmail.com" {
		t.Fatalf("user was not added")
	}

	if err := AddNewSession(ssid, "otheremail"); err == nil {
		t.Fatalf("double session not detected")
	}

}
