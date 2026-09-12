package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewAdminRouter(adminHandler *handlers.AdminHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/newSupport", adminHandler.NewSupport)
	r.Post("/transferTicket", adminHandler.TransferTicket)

	r.Get("/getTicket", adminHandler.GetTicketsWithStatus)

	return r
}
