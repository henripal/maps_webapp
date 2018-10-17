package users

import (
	"testing"
)

func Test_AddNewUser(t *testing.T) {
	var u = User{"daffy@gmail.com", "daffy", "duck", []byte("quack")}

	if err := AddNewUser(u); err != nil {
		t.Fatal(err)
	}

	if DbUser["daffy@gmail.com"].FirstName != "daffy" {
		t.Fatalf("user was not added")
	}

	if err := AddNewUser(u); err == nil {
		t.Fatalf("double user not detected")
	}

}
