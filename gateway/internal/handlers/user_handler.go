package handlers

import (
	"context"
	"encoding/json"
	pb "gateway/internal/protos"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req pb.NewTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	res, err := h.grpcClient.NewTicket(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}

func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req pb.NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	res, err := h.grpcClient.NewUser(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}

func (h *UserHandler) GetUserTickets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	userID := chi.URLParam(r, "userID")
	req := pb.GetUserTicketsRequest{UserId: userID}

	res, err := h.grpcClient.GetUserTickets(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}
