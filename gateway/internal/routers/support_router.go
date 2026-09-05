package routers

import (
	pb "gateway/internal/protos"

	"github.com/go-chi/chi/v5"
)

func NewSupportRouter(client pb.TicketServiceClient) *chi.Mux {

	r := chi.NewRouter()
	// r.Post("/answerTicket", handlers.EchoHandler)
	// r.Post("/closeTicket", handlers.EchoHandler)
	// r.Post("/transferTicket", handlers.EchoHandler)

	return r
}
