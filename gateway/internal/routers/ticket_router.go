package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewTicketRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/newTicket", handlers.EchoHandler)

	return r
}
