package routers

import (
	"gateway/internal/handlers"
	pb "gateway/internal/protos"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(client pb.TicketServiceClient) *chi.Mux {

	handler := handlers.NewUserHandler(client)

	r := chi.NewRouter()
	r.Post("/newTicket", handler.NewTicketHandler)

	return r
}
