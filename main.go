package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	color.Blue("Starting server on port 8080...")

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
