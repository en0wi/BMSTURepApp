package main

import (
	"BMSTURepApp/internal/config"
	"BMSTURepApp/internal/http-server/handlers/groups"
	"BMSTURepApp/internal/http-server/handlers/reservations"
	"BMSTURepApp/internal/http-server/handlers/users"
	"BMSTURepApp/storage"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()
	database, err := storage.NewDB(cfg.ConnectionString)

	if err != nil {
		fmt.Printf("Error with storage: %s", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Method("POST", "/userCreation/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.CreateUser(w, r, database)
	}))
	router.Method("GET", "/userReading/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.ReadUser(w, r, database)
	}))
	router.Method("PUT", "/userUpdate/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.UpdateUser(w, r, database)
	}))
	router.Method("DELETE", "/userDeletion/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.DeleteUser(w, r, database)
	}))
	router.Method("POST", "/groupCreation/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groups.CreateGroup(w, r, database)
	}))
	router.Method("GET", "/groupReading/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groups.ReadGroup(w, r, database)
	}))
	router.Method("PUT", "/groupsUpdate/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groups.UpdateGroup(w, r, database)
	}))
	router.Method("DELETE", "/groupDeletion/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groups.DeleteGroup(w, r, database)
	}))
	router.Method("POST", "/reservationCreation/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reservations.CreateReservation(w, r, database)
	}))
	router.Method("GET", "/reservationReading/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reservations.ReadReservation(w, r, database)
	}))
	router.Method("PUT", "/reservationUpdate/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reservations.UpdateReservation(w, r, database)
	}))
	router.Method("DELETE", "/reservationDeletion/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reservations.DeleteReservation(w, r, database)
	}))

	http.ListenAndServe(":3000", router)
}

// TODO : gitignore file
