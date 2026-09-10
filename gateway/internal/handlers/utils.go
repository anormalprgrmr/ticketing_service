package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func writeProtoJSON(w http.ResponseWriter, status int, msg proto.Message) {
	w.Header().Set("Content-Type", "application/json")

	data, err := protojson.Marshal(msg)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("gRPC RESPONSE reqID=%v response=%v", data)

	w.WriteHeader(status)
	_, _ = w.Write(data)

}

func writeGRPCError(w http.ResponseWriter, err error) {
	log.Errorf("gRPC error: %v", err)

	switch status.Code(err) {
	case codes.NotFound:
		http.Error(w, "not found", http.StatusNotFound)

	case codes.InvalidArgument:
		http.Error(w, "invalid request", http.StatusBadRequest)

	case codes.Unauthenticated:
		http.Error(w, "unauthorized", http.StatusUnauthorized)

	case codes.PermissionDenied:
		http.Error(w, "forbidden", http.StatusForbidden)

	case codes.DeadlineExceeded:
		http.Error(w, "request timeout", http.StatusGatewayTimeout)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
