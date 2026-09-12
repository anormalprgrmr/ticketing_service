package handlers

import (
	"context"
	"encoding/json"
	"errors"
	pb "gateway/internal/protos"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

type AdminHandler struct {
	grpcClient pb.TicketServiceClient
}

func NewAdminHandler(grpcClient pb.TicketServiceClient) *AdminHandler {
	return &AdminHandler{
		grpcClient: grpcClient,
	}
}

func (h *AdminHandler) NewSupport(w http.ResponseWriter, r *http.Request) {
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

func (h *AdminHandler) GetTicketsWithStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	queryStatus := r.URL.Query().Get("status")

	var status pb.TicketStatus
	switch queryStatus {
	case "opened":
		status = pb.TicketStatus_TICKET_STATUS_OPEN
	case "answered":
		status = pb.TicketStatus_TICKET_STATUS_ANSWERED
	case "closed":
		status = pb.TicketStatus_TICKET_STATUS_CLOSED
	default:
		writeGRPCError(w, errors.New("undefined ticket status"))
		return
	}

	req := pb.GetTicketsWithStatusRequest{Status: status}

	res, err := h.grpcClient.GetTicketsWithStatus(ctx, &req)
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}

func (h *AdminHandler) TransferTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("Start sending gRPC req")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req pb.TransferTicketRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.grpcClient.TransferTicket(ctx, &req)
	if err != nil {
		log.Printf("error while calling rpc: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeProtoJSON(w, http.StatusOK, res)
}
