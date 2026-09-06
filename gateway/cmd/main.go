package main

import (
	"fmt"
	"gateway/internal/config"
	pb "gateway/internal/protos"
	"gateway/internal/routers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	err := config.InitLogger()
	if err != nil {
		log.Error("cant create logger instance :", err)
	}

	config := config.NewConfig()
	log.Info("Starting App...")
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.RemoteHost, config.RemotePort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect to service :%v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	log.Println("Starting HTTP server...")
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	r.Mount("/api/v1/user", routers.NewUserRouter(client))
	r.Mount("/api/v1/support", routers.NewSupportRouter(client))

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
	if err != nil {
		log.Fatalf("error starting gateway %v", err.Error())
	}
}
