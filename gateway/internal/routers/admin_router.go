package routers

import (
	"github.com/go-chi/chi/v5"
)

func NewAdminRouter() *chi.Mux {

	r := chi.NewRouter()
	// r.Post("/answerTicket", handlers.EchoHandler)
	// r.Post("/closeTicket", handlers.EchoHandler)
	// r.Post("/transferTicket", handlers.EchoHandler)

	return r
}
