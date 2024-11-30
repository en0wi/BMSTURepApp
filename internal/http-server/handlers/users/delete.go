package users

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, database *storage.DB) resp.Response {
	var user structures.User
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return resp.Error("failed to decode request")
	}
	database.DeleteUser(user.Id)
	return resp.OK()
}
