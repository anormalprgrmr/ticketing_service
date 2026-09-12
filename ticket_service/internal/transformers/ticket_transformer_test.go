package transformers

import (
	"testing"
	"uuid"

	"ticket_service/internal/models"
	pb "ticket_service/internal/protos"
)

func TestConvertTicketStatusModelToGRPC(t *testing.T) {
	tests := []struct {
		name     string
		input    models.TicketStatus
		expected pb.TicketStatus
	}{
		{
			name:     "opened",
			input:    models.Opened,
			expected: pb.TicketStatus_TICKET_STATUS_OPEN,
		},
		{
			name:     "answered",
			input:    models.Answered,
			expected: pb.TicketStatus_TICKET_STATUS_ANSWERED,
		},
		{
			name:     "closed",
			input:    models.Closed,
			expected: pb.TicketStatus_TICKET_STATUS_CLOSED,
		},
		{
			name:     "unknown status",
			input:    models.TicketStatus("unknown"),
			expected: pb.TicketStatus_TICKET_STATUS_UNSPECIFIED,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertTicketStatusModelToGRPC(tt.input)

			if result != tt.expected {
				t.Errorf(
					"ConvertTicketStatusModelToGRPC(%v) = %v; want %v",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestConvertTicketStatusGRPCToModel(t *testing.T) {
	tests := []struct {
		name     string
		input    pb.TicketStatus
		expected models.TicketStatus
	}{
		{
			name:     "opened",
			input:    pb.TicketStatus_TICKET_STATUS_OPEN,
			expected: models.Opened,
		},
		{
			name:     "answered",
			input:    pb.TicketStatus_TICKET_STATUS_ANSWERED,
			expected: models.Answered,
		},
		{
			name:     "closed",
			input:    pb.TicketStatus_TICKET_STATUS_CLOSED,
			expected: models.Closed,
		},
		{
			name:     "unspecified defaults to opened",
			input:    pb.TicketStatus_TICKET_STATUS_UNSPECIFIED,
			expected: models.Opened,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertTicketStatusGRPCToModel(tt.input)

			if result != tt.expected {
				t.Errorf(
					"ConvertTicketStatusGRPCToModel(%v) = %v; want %v",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestTicketModelToGRPC(t *testing.T) {
	var id1 uuid.UUID
	var userID1 uuid.UUID
	var supportID1 uuid.UUID

	var id2 uuid.UUID
	var userID2 uuid.UUID
	var supportID2 uuid.UUID

	tickets := []*models.Ticket{
		{
			ID:        id1,
			UserID:    userID1,
			SupportID: &supportID1,
			Body:      "First ticket",
			Status:    models.Opened,
		},
		{
			ID:        id2,
			UserID:    userID2,
			SupportID: &supportID2,
			Body:      "Second ticket",
			Status:    models.Closed,
		},
	}

	result := TicketModelToGRPC(tickets)

	if len(result) != len(tickets) {
		t.Fatalf(
			"TicketModelToGRPC() returned %d tickets; want %d",
			len(result),
			len(tickets),
		)
	}

	for i, ticket := range tickets {
		got := result[i]

		if got == nil {
			t.Fatalf("result[%d] is nil", i)
		}

		if got.Id != ticket.ID.String() {
			t.Errorf(
				"result[%d].Id = %q; want %q",
				i,
				got.Id,
				ticket.ID.String(),
			)
		}

		if got.UserId != ticket.UserID.String() {
			t.Errorf(
				"result[%d].UserId = %q; want %q",
				i,
				got.UserId,
				ticket.UserID.String(),
			)
		}

		if got.SupportId != ticket.SupportID.String() {
			t.Errorf(
				"result[%d].SupportId = %q; want %q",
				i,
				got.SupportId,
				ticket.SupportID.String(),
			)
		}

		if got.Body != ticket.Body {
			t.Errorf(
				"result[%d].Body = %q; want %q",
				i,
				got.Body,
				ticket.Body,
			)
		}

		expectedStatus := ConvertTicketStatusModelToGRPC(ticket.Status)

		if got.Status != expectedStatus {
			t.Errorf(
				"result[%d].Status = %v; want %v",
				i,
				got.Status,
				expectedStatus,
			)
		}
	}
}

func TestTicketModelToGRPC_Empty(t *testing.T) {
	result := TicketModelToGRPC(nil)

	if len(result) != 0 {
		t.Errorf(
			"TicketModelToGRPC(nil) returned length %d; want 0",
			len(result),
		)
	}
}
