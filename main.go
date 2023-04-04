package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	Username string  `json:"username"`
	Avatar   string  `json:"avatar"`
	Profile  Profile `json:"profile"`
}

type Profile struct {
	Name         string        `json:"name"`
	Image        string        `json:"image"`
	Bio          string        `json:"bio"`
	Philosophy   string        `json:"philosophy"`
	Achievements []Achievement `json:"achievements"`
	SocialLinks  []Social      `json:"socialLinks"`
}

type Social struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}

type Achievement struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var users []User

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{username}", getUser).Methods("GET")
	r.HandleFunc("/users/{username}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

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

func init() {
	if _, err := os.Stat("users.json"); os.IsNotExist(err) {
		file, err := os.Create("users.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		emptyUsers := []User{}
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", " ")
		err = encoder.Encode(emptyUsers)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Loaded %d users from users.json\n", len(users))
}
