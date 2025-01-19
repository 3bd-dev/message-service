package docs

import (
	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/3bd-dev/wallet-service/internal/services/message"
	"github.com/3bd-dev/wallet-service/pkg/web"
	"github.com/google/uuid"
)

// swagger:route GET /api/v1/messages Messages ListMessages
// List all messages.
// responses:
//   200: ListMessagesResponse

// swagger:parameters ListMessages
type ListMessagesParamsWrapper struct {
	// SenderID is an optional query parameter to filter messages by sender.
	// in:query
	SenderID string `json:"sender_id"`
}

// swagger:response ListMessagesResponse
type ListMessagesResponseWrapper struct {
	// in:body
	Body struct {
		web.Response
		Data []models.Message `json:"data"`
	}
}

// swagger:route POST /api/v1/messages Messages CreateMessage
// Create a new message.
// responses:
//   201: CreateMessageResponse

// swagger:parameters CreateMessage
type CreateMessageParamsWrapper struct {
	// in:body
	Body message.NewMessage
}

// swagger:response CreateMessageResponse
type CreateMessageResponseWrapper struct {
	// in:body
	Body struct {
		web.Response
		Data models.Message `json:"data"`
	}
}

// swagger:route POST /api/v1/messages/{id}/reviews Messages ReviewMessage
// Create a new message review.
// responses:
//   200: ReviewMessageResponse

// swagger:parameters ReviewMessage
type ReviewMessageParamsWrapper struct {
	// ID of the message to review.
	// in:path
	// required: true
	// example: 123e4567-e89b-12d3-a456-426614174000
	ID uuid.UUID `json:"id"`

	// in:body
	Body message.NewMessageReview
}

// swagger:response ReviewMessageResponse
type ReviewMessageResponseWrapper struct {
	// in:body
	Body struct {
		web.Response
		Data models.MessageReview `json:"data"`
	}
}
