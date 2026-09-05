package handlers

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	pb "gateway/internal/protos"
	"net/http"
	"time"
)

type UserHandler struct {
	grpcClient pb.TicketServiceClient
}

func NewUserHandler(grpcClient pb.TicketServiceClient) *UserHandler {
	return &UserHandler{
		grpcClient: grpcClient,
	}
}

func (h *UserHandler) NewTicketHandler(w http.ResponseWriter, r *http.Request) {

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
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	var buf bytes.Buffer

	err = binary.Write(&buf, binary.BigEndian, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	w.Write(buf.Bytes())

}
