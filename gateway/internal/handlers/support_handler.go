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

type SupportHandler struct {
	grpcClient pb.TicketServiceClient
}

func NewSupportHandler(grpcClient pb.TicketServiceClient) *SupportHandler {
	return &SupportHandler{
		grpcClient: grpcClient,
	}
}

func (h *SupportHandler) NewSupport(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.NewSupportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.NewSupport(ctx, &req)
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

func (h *SupportHandler) GetSupportTickets(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	supportID := chi.URLParam(r, "supportID")
	req := pb.GetSupportTicketsRequest{SupportId: supportID}

	res, err := h.grpcClient.GetSupportTickets(ctx, &req)
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

func (h *SupportHandler) CloseTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.CloseTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.CloseTicket(ctx, &req)
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

func (h *SupportHandler) AnswerTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.AnswerTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.AnswerTicket(ctx, &req)
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
