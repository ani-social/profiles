package main

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	color.Green("GET request received for all users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	color.Cyan("POST request received to create a new user")
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, user)
	saveUsers()
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	color.Yellow("GET request received for a specific user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	for _, user := range users {
		if user.Username == username {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	color.Magenta("PUT request received to update a user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	for i, user := range users {
		if user.Username == username {
			var updatedUser User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			users[i] = updatedUser
			saveUsers()
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.NotFound(w, r)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	color.Red("DELETE request received to delete a user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	for i, user := range users {
		if user.Username == username {
			users = append(users[:i], users[i+1:]...)
			saveUsers()
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func saveUsers() {
	dataFile, err := os.Create("users.json")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer dataFile.Close()
	encoder := json.NewEncoder(dataFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(users)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Green("Users saved successfully")
}
