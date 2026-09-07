package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewAdminRouter(adminHandler *handlers.AdminHandler) *chi.Mux {

	r := chi.NewRouter()
	r.Post("/newSupport", adminHandler.NewSupport)
	// r.Post("/closeTicket", handlers.EchoHandler)
	// r.Post("/transferTicket", handlers.EchoHandler)

	return r
}
