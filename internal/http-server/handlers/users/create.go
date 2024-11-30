package create

import (
	resp "BMSTURepApp/internal/lib/response"
	"github.com/go-chi/render"
	"net/http"
)

type User struct {
	firstName   string
	surname     string
	lastName    string
	groupID     string
	phoneNumber string
	telegramTag string
	vkLink      string
}

type Response struct {
	resp.Response
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		err := render.DecodeJSON(r.Body, &user)

		if err != nil {
			// log = ... надо дописать логгер
			render.JSON(w, r, resp.Error("Failed to decode request"))
			return
		}

	}
}
