package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var tests = []struct {
	un       string
	pw       string
	expected int
}{
	{"bob", "hunter2", http.StatusOK},
	{"alfred", "passWord", http.StatusNotFound},
	{"alice", "wrongPassword", http.StatusForbidden},
}

func TestHMACToken(t *testing.T) {
	userMap = make(map[string]User)
	config, _ = LoadConfig("config.json")
	for _, test := range tests {
		//Test each user
		path := fmt.Sprintf("/session/create")

		req, err := http.NewRequest("GET", path, nil)
		req.SetBasicAuth(test.un, test.pw)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/session/create", GenerateHMAC)
		router.ServeHTTP(rr, req)

		if test.un == "bob" && rr.Code != http.StatusOK {
			t.Errorf("\nExpected: %d, but got: %d, Data: %s", test.expected, rr.Code, test.un)
		}

		if test.un == "alfred" && rr.Code != http.StatusNotFound {
			t.Errorf("\nExpected: %d, but got: %d, Data: %s", test.expected, rr.Code, test.un)
		}

		if test.un == "alice" && rr.Code != http.StatusForbidden {
			t.Errorf("\nExpected: %d, but got: %d, Data: %s", test.expected, rr.Code, test.un)
		}
	}
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//TODO
//Edge case tests, null cases, bad data

//TODO
//More tests for HMAC

//TODO
//Reverse proxy tests
