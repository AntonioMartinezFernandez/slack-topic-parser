package http_server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	// Http Server and Router
	fmt.Println("Starting server at localhost:3000")
	gorillaMux := mux.NewRouter()

	gorillaMux.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	gorillaMux.HandleFunc("/addUser", addUserHandler).Methods("POST")
	gorillaMux.HandleFunc("/addFirstUser", addFirstUserHandler).Methods("POST")
	gorillaMux.HandleFunc("/removeUser", removeUserHandler).Methods("POST")

	srv := &http.Server{
		Handler: gorillaMux,
		Addr:    "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
