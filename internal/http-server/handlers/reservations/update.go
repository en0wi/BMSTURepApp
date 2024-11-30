package groups

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func UpdateGroup(w http.ResponseWriter, r *http.Request, database *storage.DB) {
	var group structures.Group

	err := render.DecodeJSON(r.Body, &group)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.UpdateGroupinfo(group)

	data, err := json.Marshal(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
