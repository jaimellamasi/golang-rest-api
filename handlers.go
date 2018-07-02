package api

import (
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User struct represents a user in the system with name and address
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Job  string `json:"job"`
}

var users = []User{
	User{1, "Maxim", "Ex ministro"},
	User{2, "Julen", "Ex seleccionador"},
	User{3, "Mariano", "Ex presidente"},
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, 200, users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		utils.WriteError(w, 400, "user ID must be a number")
	}
	for _, user := range users {
		if user.ID == userID {
			utils.WriteJSON(w, 200, user)
			return
		}
	}
	utils.WriteError(w, 404, "User not found")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		utils.WriteError(w, 400, err.Error())
		return
	}
	defer r.Body.Close()
	lastID := users[len(users)-1].ID
	user.ID = lastID + 1
	users = append(users, user)
	utils.WriteJSON(w, 200, user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		utils.WriteError(w, 400, "user ID must be a number")
	}
	decoder := json.NewDecoder(r.Body)
	var userUp User
	err = decoder.Decode(&userUp)
	if err != nil {
		utils.WriteError(w, 400, err.Error())
		return
	}
	defer r.Body.Close()
	for _, user := range users {
		if user.ID == userID {
			user.Name = userUp.Name
			user.Job = userUp.Job
			utils.WriteJSON(w, 200, user)
			return
		}
	}
	utils.WriteError(w, 404, "User not found")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		utils.WriteError(w, 400, "user ID must be a number")
	}
	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			utils.WriteJSON(w, 200, nil)
			return
		}
	}
	utils.WriteError(w, 404, "User not found")
}
