package message

import (
	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/google/uuid"
)

type NewMessage struct {
	Content   string    `json:"content" validate:"required,min=5,max=100"`
	Recipient string    `json:"recipient" validate:"required"`
	SenderID  uuid.UUID `json:"sender_id" validate:"required,uuid"`
}

type NewMessageReview struct {
	Action     models.MessageReviewAction `json:"action" validate:"required,oneof=approved rejected"`
	ReviewerID uuid.UUID                  `json:"reviewer_id" validate:"required,uuid"`
}
