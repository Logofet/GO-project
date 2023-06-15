package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *UserModel) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := u.getAllUsers()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while getting all users", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%v", users)
	case http.MethodPost:
		var user User
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while parsing form", http.StatusBadRequest)
			return
		}

		user.Name = r.FormValue("name")
		user.Email = r.FormValue("email")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while rounding age", http.StatusBadRequest)
			return
		}
		user.Age = age

		err = u.addUser(user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while adding new user", http.StatusInternalServerError)
			return
		}
	}
}

func (u *UserModel) SingleUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	switch r.Method {
	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while decoding request body", http.StatusBadRequest)
		}

		err = u.updateUser(user, id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while updating user", http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		err := u.deleteUser(id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while deleting user", http.StatusInternalServerError)
			return
		}
	}
}
