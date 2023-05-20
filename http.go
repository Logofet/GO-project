package main

import (
	"Projekat/server"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func httpServer(db *sql.DB) {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.Handler1)

	router.HandleFunc("/users/{id}", server.Handler2)

	log.Println("Listening on port 6969...")
	if err := http.ListenAndServe(":6969", router); err != nil {
		log.Fatal(err)
	}

}
