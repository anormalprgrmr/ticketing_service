package handlers

import (
	"context"
	"encoding/json"
	pb "gateway/internal/protos"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"
)

type UserHandler struct {
	grpcClient pb.TicketServiceClient
}

func NewUserHandler(grpcClient pb.TicketServiceClient) *UserHandler {
	return &UserHandler{
		grpcClient: grpcClient,
	}
}

func (h *UserHandler) NewTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.NewTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.NewTicket(ctx, &req)
	if err != nil {
		log.Printf("error while calling rpc: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	log.Printf("response is => %v", res)

	resByte, err := protojson.Marshal(res)
	if err != nil {
		log.Printf("error while converting proto to json : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	w.Write(resByte)

}

func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.NewUser(ctx, &req)
	if err != nil {
		log.Printf("error while calling rpc: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	log.Printf("response is => %v", res)

	resByte, err := protojson.Marshal(res)
	if err != nil {
		log.Printf("error while converting proto to json : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	w.Write(resByte)

}

func (h *UserHandler) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userID := chi.URLParam(r, "userID")
	req := pb.GetUserTicketsRequest{UserId: userID}

	res, err := h.grpcClient.GetUserTickets(ctx, &req)
	if err != nil {
		log.Printf("error while calling rpc: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	log.Printf("response is => %v", res)

	resByte, err := protojson.Marshal(res)
	if err != nil {
		log.Printf("error while converting proto to json : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	w.Write(resByte)

}
