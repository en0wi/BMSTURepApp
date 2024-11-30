package reservations

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	resp.Response
}

func CreateReservation(w http.ResponseWriter, r *http.Request, database *storage.DB) {
	var reservation structures.Reservation

	err := render.DecodeJSON(r.Body, &reservation)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return
	}
	database.CreateReserv(reservation)
}
