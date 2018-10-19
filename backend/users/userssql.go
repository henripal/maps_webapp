package users

import (
	"database/sql"
	"log"
	"os"

	// driver
	_ "github.com/lib/pq"
)

// DB is a pointer towards the Users database objedt
// it is initialized by InitializeDB
var DB *sql.DB

// InitializeDB instanciates and intializes the User database into the
// DBU global variable
func InitializeDB() error {
	DBIP := os.Getenv("DBIP")
	var connStr string
	if DBIP == "" {
		connStr = "user=henripal dbname=webapp sslmode=disable"
	} else {
		connStr = "user=crud dbname=webapp sslmode=disable password=smallpers0n host=" + DBIP
	}
	var err error
	DB, err = sql.Open("postgres", connStr)

	return err
}

// GetUser returns the User structure from the User database
// corresponding to the email email
func GetUser(email string) (User, error) {
	var u User
	row := DB.QueryRow("SELECT * FROM USERS WHERE email=$1", email)
	err := row.Scan(&u.Email, &u.FirstName, &u.LastName, &u.Password)
	return u, err
}

// DeleteUser  deletes the user with key email
func DeleteUser(email string) error {
	var u User
	row := DB.QueryRow("SELECT * FROM USERS WHERE email=$1", email)
	err := row.Scan(&u.Email, &u.FirstName, &u.LastName, &u.Password)
	if err != nil {
		log.Fatalln("Could not delete user.")
	}
	sqlStatement := `DELETE FROM USERS WHERE email=$1`
	_, err = DB.Exec(sqlStatement, email)
	return err
}

// AddUser adds User u to the database
func addUser(u User) error {
	sqlStatement := `
		INSERT INTO USERS (email, first, last, password)
		VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(sqlStatement, u.Email, u.FirstName, u.LastName, u.Password)
	return err
}
