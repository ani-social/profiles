// json.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

	// Update number of loaded JSON items when a user is added or deleted
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "DELETE" {
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
	})
}
