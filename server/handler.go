package server

import (
	"Projekat/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Handler1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := database.Db.Query("SELECT * FROM users")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while getting users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error while getting users", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		log.Println(users)
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

		_, err = database.Db.Exec("INSERT INTO users (id, name, email, age) VALUES (0, ?, ?, ?)", user.Name, user.Email, user.Age)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while adding new user", http.StatusInternalServerError)
			return
		}

	}
}
func Handler2(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	switch r.Method {
	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while decoding request body", http.StatusBadRequest)
		}

		_, err = database.Db.Exec("UPDATE users SET name=?, email=?, age=? WHERE id=?", user.Name, user.Email, user.Age, id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while updating user "+id, http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		_, err := database.Db.Exec(fmt.Sprintf("DELETE FROM users WHERE ID = %s", id))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while deleting user "+id, http.StatusInternalServerError)
			return
		}
	}
}
