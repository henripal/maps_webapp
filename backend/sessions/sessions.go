package sessions

import (
	"errors"
)

// DbSession is the stub of the session database
var DbSession = map[string]string{}

// AddNewSession creates a new session in database
func AddNewSession(uuid string, email string) error {
	if _, uuidAlreadyInDb := DbSession[uuid]; uuidAlreadyInDb {
		return errors.New("uuid is already in database")
	}
	DbSession[uuid] = email
	return nil
}
