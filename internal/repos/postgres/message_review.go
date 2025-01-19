package postgres

import (
	"context"

	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/3bd-dev/wallet-service/pkg/database"
	"github.com/google/uuid"
)

type MessageReviewRepo struct {
	db database.IDatabase
}

// NewMessageReviewRepo creates a new instance of MessageReviewRepo.
func NewMessageReviewRepo(db database.IDatabase) *MessageReviewRepo {
	return &MessageReviewRepo{db: db}
}

// Create a new message review in the database.
func (r *MessageReviewRepo) Create(ctx context.Context, review *models.MessageReview) error {
	return r.db.WithContext(ctx).Create(review).Error
}

// HasUserReviewed checks if a user has already reviewed a message.
func (r *MessageReviewRepo) HasUserReviewed(ctx context.Context, messageID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.MessageReview{}).
		Where("message_id = ? AND reviewer_id = ?", messageID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CountApprovals returns the number of approvals for a message.
func (r *MessageReviewRepo) CountApprovals(ctx context.Context, messageID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.MessageReview{}).
		Where("message_id = ? AND action = ?", messageID, models.MessageReviewActionApproved).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// HasRejection checks if a message has been rejected by any reviewer.
func (r *MessageReviewRepo) HasRejection(ctx context.Context, messageID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.MessageReview{}).
		Where("message_id = ? AND action = ?", messageID, models.MessageReviewActionRejected).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
