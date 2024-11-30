package reservations

import (
	structures "BMSTURepApp/internal/domain"
	resp "BMSTURepApp/internal/lib/response"
	"BMSTURepApp/storage"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteReservation(w http.ResponseWriter, r *http.Request, database *storage.DB) resp.Response {
	var reservation structures.Reservation
	err := render.DecodeJSON(r.Body, &reservation)
	if err != nil {
		render.JSON(w, r, resp.Error("Failed to decode request"))
		return resp.Error("failed to decode request")
	}
	database.DeleteReserv(reservation.Id)
	return resp.OK()
}
