package groups

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage/db"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func UpdateGroup(w http.ResponseWriter, r *http.Request, database *db.DB) {
	var group structures.Group

	err := render.DecodeJSON(r.Body, &group)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.Update_Groupinfo(group)

	data, err := json.Marshal(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
