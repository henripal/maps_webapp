package users

import (
	"log"
	"testing"
)

func init() {
	if err := InitializeDBUsers(); err != nil {
		log.Fatalln(err)
	}
}
func Test_AddNewUser(t *testing.T) {
	var u = User{"daffy@gmail.com", "daffy", "duck", []byte("quack")}

	if err := AddNewUser(u); err != nil {
		t.Fatal(err)
	}

	u, err := GetUser("daffy@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	if u.FirstName != "daffy" {
		t.Fatalf("user was not added")
	}

	if err := AddNewUser(u); err == nil {
		t.Fatalf("double user not detected")
	}

	err = DeleteUser("daffy@gmail.com")
	if err != nil {
		t.Fatalf("Unable to delete test user")
	}

}
