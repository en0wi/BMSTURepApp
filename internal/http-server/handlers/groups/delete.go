package groups

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage/db"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteGroup(w http.ResponseWriter, r *http.Request, database *db.DB) resp.Response {
	var group structures.Group
	err := render.DecodeJSON(r.Body, &group)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return resp.Error("failed to decode request")
	}
	database.Delete_Group(group.Id)
	return resp.OK()
}
