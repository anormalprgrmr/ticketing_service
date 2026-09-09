package routers

import (
	"gateway/internal/handlers"
	pb "gateway/internal/protos"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(client pb.TicketServiceClient) *chi.Mux {
	userHandler := handlers.NewUserHandler(client)
	supportHandler := handlers.NewSupportHandler(client)
	adminHandler := handlers.NewAdminHandler(client)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.CleanPath)

	r.Use(middleware.Logger)

	r.Mount("/api/user", NewUserRouter(userHandler))
	r.Mount("/api/support", NewSupportRouter(supportHandler))
	r.Mount("/api/admin", NewAdminRouter(adminHandler))

	return r
}
