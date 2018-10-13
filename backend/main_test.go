package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_data(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	data(res, req)

	exp := "WHAT IS THE TEMPLATE????"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}
