package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type User struct {
	userPermissions string
	firstName       string
	surname         string
	lastName        string
	groupID         string
	phoneNumber     string
	telegramTag     string
	vkLink          string
}

type Request struct {
	RequestType        string
	RequestPermissions string
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "telegramTag")
	// user ...
	user, err := database.GetUser(userID)

	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("Error fetching user %s: %v", userID, err)))
		return
	}

	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("User not found"))
		return
	}
	return user.id
}

func main() {
	mainRouter := chi.NewRouter()
	mainRouter.Get("/users/{userID}", getUser)
	http.ListenAndServe(":3000", mainRouter)
}

//test
