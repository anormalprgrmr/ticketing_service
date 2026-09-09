package routers

import (
	"gateway/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewSupportRouter(supportHandler *handlers.SupportHandler) *chi.Mux {

	r := chi.NewRouter()

	r.Post("/answerTicket", supportHandler.NewSupport)
	r.Post("/closeTicket", supportHandler.CloseTicket)
	r.Post("/answerTicket", supportHandler.AnswerTicket)

	r.Get("/{userID}/myTickets", supportHandler.GetSupportTickets)

	return r
}
