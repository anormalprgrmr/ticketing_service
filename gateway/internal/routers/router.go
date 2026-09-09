package routers

import (
	"gateway/internal/handlers"
	pb "gateway/internal/protos"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(client pb.TicketServiceClient) *chi.Mux {
	userHandler := handlers.NewUserHandler(client)
	supportHandler := handlers.NewSupportHandler(client)
	adminHandler := handlers.NewAdminHandler(client)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	r.Mount("/api/user", NewUserRouter(userHandler))
	r.Mount("/api/support", NewSupportRouter(supportHandler))
	r.Mount("/api/admin", NewAdminRouter(adminHandler))

	return r
}
