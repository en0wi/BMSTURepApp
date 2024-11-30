package users

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func ReadUser(w http.ResponseWriter, r *http.Request, database *storage.DB) {
	var user structures.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.ReadUserinfo(user.Id)

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
