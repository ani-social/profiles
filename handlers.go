// handlers.go
package main

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary		Get all users
// @Description	Retrieve a list of all users
// @ID				get-users
// @Produce		json
// @Success		200	{array}	User
// @Router			/users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	color.Green("GET request received for all users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// @Summary		Create a new user
// @Description	Create a new user with the given data
// @ID				create-user
// @Accept			json
// @Produce		json
// @Param			user	body		User	true	"User to be created"
// @Success		200		{object}	User
// @Router			/users [post]
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

// @Summary		Get a specific user
// @Description	Retrieve a user by their username
// @ID				get-user
// @Produce		json
// @Param			username	path		string	true	"Username of the user to be fetched"
// @Success		200			{object}	User
// @Failure		404			"User not found"
// @Router			/users/{username} [get]
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

// @Summary		Update a user
// @Description	Update a user's data by their username
// @ID				update-user
// @Accept			json
// @Produce		json
// @Param			username	path		string	true	"Username of the user to be updated"
// @Param			updatedUser	body		User	true	"Updated user data"
// @Success		200			{object}	User
// @Failure		404			"User not found"
// @Router			/users/{username} [put]
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

			// Merge the updated fields with the existing user data
			if updatedUser.Username != "" {
				user.Username = updatedUser.Username
			}
			if updatedUser.Avatar != "" {
				user.Avatar = updatedUser.Avatar
			}
			profile := &user.Profile
			updatedProfile := &updatedUser.Profile
			if updatedProfile.Name != "" {
				profile.Name = updatedProfile.Name
			}
			if updatedProfile.Image != "" {
				profile.Image = updatedProfile.Image
			}
			if updatedProfile.Bio != "" {
				profile.Bio = updatedProfile.Bio
			}
			if updatedProfile.Philosophy != "" {
				profile.Philosophy = updatedProfile.Philosophy
			}

			users[i] = user
			saveUsers()
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

// @Summary		Delete a user
// @Description	Delete a user by their username
// @ID				delete-user
// @Produce		json
// @Param			username	path		string	true	"Username of the user to be deleted"
// @Success		200			{object}	User
// @Failure		404			"User not found"
// @Router			/users/{username} [delete]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	color.Red("DELETE request received to delete a user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]

	if username != "Space Mommy" || username != "Space%20Mommy" {

		for i, user := range users {
			if user.Username == username {
				users = append(users[:i], users[i+1:]...)
				saveUsers()
				json.NewEncoder(w).Encode(user)
				return
			}
		}
	}

	http.NotFound(w, r)
}

func cleanUpEmptyProfiles() {
	newUsers := []User{}
	for _, user := range users {
		if user.Username != "" && user.Avatar != "" &&
			user.Profile.Name != "" && user.Profile.Image != "" &&
			user.Profile.Bio != "" && user.Profile.Philosophy != "" {
			newUsers = append(newUsers, user)
		}
	}
	users = newUsers
}
