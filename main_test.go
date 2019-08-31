package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2,3)
	assert.NotNil(t, total, "The total should not be nil")
	assert.Equal(t, 5, total, "The total should be 5")
}

func TestSubtract(t *testing.T) {
	total := Subtract(8,4)
	assert.NotNil(t, total, "The total should not be nil")
	assert.Equal(t, 4, total, "The total should be 4")
}

func TestFactorial(t *testing.T) {
	total := Factorial(2 )
	assert.NotNil(t, total, "The total should not be nil")
	assert.Equal(t, 2, total, "The total should be 2")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}
func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response expected" )
	assert.Equal(t, "Hello World", response.Body.String(), "200 response expected" )
}


