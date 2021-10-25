package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/api/hello_world", helloWorld)
	r.HandleFunc("/api/hello_user/{user}", helloUser)
	r.HandleFunc("/api/error", errorAnswer)

	http.ListenAndServe(":8080", r)
}
