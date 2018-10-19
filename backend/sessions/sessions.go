package sessions

import (
	"errors"
	"log"
	"webapp_template/backend/users"
)

// AddNewSession creates a new session in database
func AddNewSession(uuid string, email string) error {
	if _, err := GetEmailFromSession(uuid); err == nil {
		return errors.New("uuid is already in database")
	}
	err := addSession(uuid, email)
	return err
}

// GetEmailFromSession returns the Email rom the Sessions database
// corresponding to the given uuid
func GetEmailFromSession(uuid string) (string, error) {
	var email string

	row := users.DB.QueryRow("SELECT email FROM SESSIONS WHERE uuid=$1", uuid)
	err := row.Scan(&email)
	return email, err
}

// DeleteSession  deletes the session with key uuid
func DeleteSession(uuid string) error {
	_, err := GetEmailFromSession(uuid)
	if err != nil {
		log.Fatalln("Could not delete user.")
	}
	sqlStatement := `DELETE FROM SESSIONS WHERE uuid=$1`
	_, err = users.DB.Exec(sqlStatement, uuid)
	return err
}

// DeleteSessionFromEmail  deletes the session with key uuid
func DeleteSessionFromEmail(email string) error {
	sqlStatement := `DELETE FROM SESSIONS WHERE email=$1`
	_, err := users.DB.Exec(sqlStatement, email)
	return err
}

// addSession adds a session to the database
func addSession(uuid string, email string) error {
	sqlStatement := `
		INSERT INTO SESSIONS (uuid, email)
		VALUES ($1, $2)`

	_, err := users.DB.Exec(sqlStatement, uuid, email)
	return err
}
