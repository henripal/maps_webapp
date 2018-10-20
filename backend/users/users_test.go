package users

import (
	"log"
	"testing"
)

func init() {
	if err := InitializeDB(); err != nil {
		log.Fatalln(err)
	}
}
func Test_AddNewUser(t *testing.T) {
	email := "not@areal.email"
	firstName := "not"
	lastName := "areal"
	password := "quack"
	var u = User{email, firstName, lastName, []byte(password)}

	if err := AddNewUser(u); err != nil {
		t.Fatal(err)
	}

	u, err := GetUser(email)
	if err != nil {
		t.Fatal(err)
	}
	if u.FirstName != firstName {
		t.Fatalf("user was not added")
	}

	if err := AddNewUser(u); err == nil {
		t.Fatalf("double user not detected")
	}

	err = DeleteUser(email)
	if err != nil {
		t.Fatalf("Unable to delete test user")
	}

}
