package users

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage/db"
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	resp.Response
}

func CreateUser(w http.ResponseWriter, r *http.Request, database *db.DB) {
	var user structures.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.Create_User(user)
}
