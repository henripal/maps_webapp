package usersapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Users(t *testing.T) {
	_, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	_ = httptest.NewRecorder()

}
