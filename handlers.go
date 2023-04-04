package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(err)
		return
	}
	defer dataFile.Close()
	encoder := json.NewEncoder(dataFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
		return
	}
}
