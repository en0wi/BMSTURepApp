package users

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage/db"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request, database *db.DB) {
	var user structures.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.Update_Userinfo(user)

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
