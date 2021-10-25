package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	SuccessResponse(w, "Hello, world")
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := fmt.Sprintf("Hello, %s", params["user"])
	SuccessResponse(w, message)
}
func errorAnswer(w http.ResponseWriter, r *http.Request) {

	err := errors.New("some error")
	ErrorResponse(w, err)
}
