package groups

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

func CreateGroup(w http.ResponseWriter, r *http.Request, database *db.DB) {
	var group structures.Group

	err := render.DecodeJSON(r.Body, &group)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.Create_Group(group)
}
