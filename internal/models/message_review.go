package models

import (
	"time"

	"github.com/google/uuid"
)

// MessageReviewAction represents the possible actions for a message review.
type MessageReviewAction string

const (
	MessageReviewActionApproved MessageReviewAction = "approved"
	MessageReviewActionRejected MessageReviewAction = "rejected"
)

// MessageReview represents a review (approval or rejection) of a message by a checker.
type MessageReview struct {
	ID         uuid.UUID           `json:"id"`          // Unique identifier for the review
	MessageID  uuid.UUID           `json:"message_id"`  // ID of the message being reviewed
	ReviewerID uuid.UUID           `json:"reviewer_id"` // ID of the user who performed the review
	Action     MessageReviewAction `json:"action"`      // Action taken (approved or rejected)
	CreatedAt  time.Time           `json:"created_at"`  // Timestamp when the review was created
}
