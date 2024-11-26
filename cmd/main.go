package main

import (
	"BMSTURepApp/internal/config"
	"BMSTURepApp/storage/postgre"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

type Request struct {
	RequestType        string
	RequestPermissions string
}

//func getUser(w http.ResponseWriter, r *http.Request) {
//	userID := chi.URLParam(r, "telegramTag")
//	// user ...
//	user, err :
//
//	if err != nil {
//		w.WriteHeader(422)
//		w.Write([]byte(fmt.Sprintf("Error fetching user %s: %v", userID, err)))
//		return
//	}
//
//	if user == nil {
//		w.WriteHeader(404)
//		w.Write([]byte("User not found"))
//		return
//	}
//}

func main() {
	cfg := config.Load()
	fmt.Println(cfg)

	storage, err := postgre.New(cfg.StoragePath)
	if err != nil {
		fmt.Printf("Error with storage: %s", err)
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()
	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	http.ListenAndServe(":3000", router)
}

//test
