package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewSupportRouter(supportHandler *handlers.SupportHandler) *chi.Mux {

	r := chi.NewRouter()
	r.Post("/answerTicket", supportHandler.NewSupport)
	// r.Post("/closeTicket", handlers.EchoHandler)
	// r.Post("/transferTicket", handlers.EchoHandler)

	return r
}
