package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserfound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("We were not expecting a user with ID 0.")
	}
	if err == nil {
		t.Error("We were expecting an error when user ID is 0.")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting 404 erro when user is not found.")
	}
}
