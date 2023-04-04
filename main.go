package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"net/http"
	_ "profiles/docs" // Replace this with the path to your generated docs
)

// @title User Profile API
// @version 1.0
// @description This is a simple API for managing user profiles.
// @termsOfService http://anij.bytecats.codes/tos/
// @contact.name API Support
// @contact.url http://www.anij.bytecats.codes/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	color.Blue("Starting server on port 8080...")

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
