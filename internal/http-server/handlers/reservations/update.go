package reservations

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func UpdateReservation(w http.ResponseWriter, r *http.Request, database *storage.DB) {
	var reservation structures.Reservation

	err := render.DecodeJSON(r.Body, &reservation)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.UpdateReservInfo(reservation)

	data, err := json.Marshal(reservation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
