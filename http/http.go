package http

import (
	"Projekat/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpServer(userModel *models.UserModel) error {
	router := mux.NewRouter()

	router.HandleFunc("/users", userModel.UsersHandler)
	router.HandleFunc("/users/{id}", userModel.SingleUserHandler)

	log.Println("Listening on port 6969...")
	if err := http.ListenAndServe(":6969", router); err != nil {
		return err
	}

	return nil
}
