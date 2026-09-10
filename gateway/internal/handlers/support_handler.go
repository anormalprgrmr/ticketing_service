package handlers

import (
	"context"
	"encoding/json"
	pb "gateway/internal/protos"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req pb.NewSupportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	res, err := h.grpcClient.NewSupport(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)

}

func (h *SupportHandler) GetSupportTickets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	supportID := chi.URLParam(r, "supportID")
	req := pb.GetSupportTicketsRequest{SupportId: supportID}

	res, err := h.grpcClient.GetSupportTickets(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}

func (h *SupportHandler) CloseTicket(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req pb.CloseTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	res, err := h.grpcClient.CloseTicket(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)

}

func (h *SupportHandler) AnswerTicket(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req pb.AnswerTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	res, err := h.grpcClient.AnswerTicket(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)

}
