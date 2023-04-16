package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:mysqlmysql@tcp(127.0.0.1:3306)/mojabaza")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	for {
		router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				rows, err := db.Query("SELECT * FROM users")
				if err != nil {
					log.Println(err)
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				defer rows.Close()

				var users []User
				for rows.Next() {
					var user User
					err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
					if err != nil {
						log.Println(err)
						http.Error(w, "Internal server error", http.StatusInternalServerError)
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
					http.Error(w, "Bad request", http.StatusBadRequest)
					return
				}

				user.Name = r.FormValue("name")
				user.Email = r.FormValue("email")
				age, err := strconv.Atoi(r.FormValue("age"))
				if err != nil {
					log.Println(err)
					http.Error(w, "Bad request", http.StatusBadRequest)
					return
				}
				user.Age = age

				_, err = db.Exec("INSERT INTO users (id, name, email, age) VALUES (0, ?, ?, ?)", user.Name, user.Email, user.Age)
				if err != nil {
					log.Println(err)
					http.Error(w, "Error", http.StatusConflict)
					return
				}
			}
		})

		router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := mux.Vars(r)["id"]
			switch r.Method {
			case http.MethodPut:
				var user User
				err := json.NewDecoder(r.Body).Decode(&user)
				if err != nil {
					log.Println("Error")
					http.Error(w, "Error", 404)
				}

				_, err = db.Exec("UPDATE users SET name=?, email=?, age=? WHERE id=?", user.Name, user.Email, user.Age, id)
				if err != nil {
					log.Println(err)
					http.Error(w, "Greska u update", http.StatusConflict)
					return
				}
			case http.MethodDelete:
				_, err = db.Exec(fmt.Sprintf("DELETE FROM users WHERE ID = %s", id))
				if err != nil {
					log.Println(err)
					http.Error(w, "Error", http.StatusConflict)
					return
				}
			}
		})

		log.Println("Listening on port 6969...")
		if err := http.ListenAndServe(":6969", router); err != nil {
			log.Fatal(err)
		}
	}

}
