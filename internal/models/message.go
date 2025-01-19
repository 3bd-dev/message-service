package models

import (
	"time"

	"github.com/google/uuid"
)

// MessageStatus represents the possible statuses of a message.
type MessageStatus string

const (
	MessageStatusPending  MessageStatus = "pending"
	MessageStatusApproved MessageStatus = "approved"
	MessageStatusRejected MessageStatus = "rejected"
)

// Message represents a message created by a sender.
type Message struct {
	ID                uuid.UUID       `json:"id"`                                   // Unique identifier for the message
	Content           string          `json:"content"`                              // Content of the message
	Status            MessageStatus   `json:"status"`                               // Status of the message (pending, approved, rejected)
	Recipient         string          `json:"recipient"`                            // Recipient of the message (e.g., email)
	SenderID          uuid.UUID       `json:"sender_id"`                            // ID of the user who sent the message
	ApprovalsRequired int             `json:"approvals_required"`                   // Number of approvals required
	CreatedAt         time.Time       `json:"created_at"`                           // Timestamp when the message was created
	UpdatedAt         time.Time       `json:"updated_at"`                           // Timestamp when the message was last updated
	Reviews           []MessageReview `gorm:"foreignKey:message_id" json:"reviews"` // List of MessageReviews
}
