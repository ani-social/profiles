package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
