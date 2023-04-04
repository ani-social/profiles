package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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
// @host pipebomb.bytecats.codes
// @BaseURL pipebomb.bytecats.codes
// @BasePath /
// @schemes https
func main() {
	color.Blue("Starting server on port 8080...")

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{username}", deleteUser).Methods("DELETE")
	// add the docs httpSwagger handler
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	// Adding CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
