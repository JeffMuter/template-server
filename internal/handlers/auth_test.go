package handlers

import "testing"

func TestLoginFormHandler(t *testing.T) {
	// create a struct that we can add data to in an array of said struct
	// make a loop to login, using the values of each of the structs,
	// set test conditions for each
	// 1. Sign in with knowingly bad credentials
	// 2. Sign in with knowingly good credentials
	// 3. Sign in with empty credentials
	// 4. Sign in with variations of partially good credentials.
}

func TestResigterUserHandler(t *testing.T) {
	// create dummy data, loop through each, attempting to register
	// for each. Trying valid sends, partially correct, and empty attempts.
}
