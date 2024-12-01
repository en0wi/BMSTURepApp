package users

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage/db"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, database *db.DB) resp.Response {
	var user structures.User
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return resp.Error("failed to decode request")
	}
	database.Delete_User(user.Id)
	return resp.OK()
}
