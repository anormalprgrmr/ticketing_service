package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(userHandler *handlers.UserHandler) *chi.Mux {

	r := chi.NewRouter()

	r.Post("/newUser", userHandler.NewUser)
	r.Post("/newTicket", userHandler.NewTicket)

	r.Get("/{userID}/myTickets", userHandler.NewTicket)

	return r
}
