package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/3bd-dev/wallet-service/internal/services/message"
	"github.com/3bd-dev/wallet-service/pkg/database"
	"github.com/3bd-dev/wallet-service/pkg/errs"
	"github.com/3bd-dev/wallet-service/pkg/page"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db database.IDatabase
}

// NewMessageRepo creates a new instance of MessageRepo.
func NewMessageRepo(db database.IDatabase) *MessageRepo {
	return &MessageRepo{db: db}
}

// Create a new msg record in the database.
func (r *MessageRepo) Create(ctx context.Context, msg *models.Message) error {
	return r.db.WithContext(ctx).Create(msg).Error
}

// GetByID get message from database
func (r *MessageRepo) QueryByID(ctx context.Context, id uuid.UUID) (*models.Message, error) {
	var msg models.Message
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&msg).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.NotFound, fmt.Errorf("message with ID %s not found", id))
		}
		return nil, err
	}

	return &msg, nil
}

// Update a message in the database
func (r *MessageRepo) Update(ctx context.Context, msg *models.Message) error {
	return r.db.WithContext(ctx).Updates(msg).Error
}

// Query to list all messages in the database with filter
func (r *MessageRepo) Query(ctx context.Context, filter message.QueryFilter, page page.Page) ([]models.Message, int64, error) {
	var msgs []models.Message
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Message{}).Preload("Reviews")

	if filter.SenderID != nil {
		query = query.Where("sender_id = ?", filter.SenderID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(page.Offset()).Limit(page.RowsPerPage()).Find(&msgs).Error; err != nil {
		return nil, 0, err
	}

	return msgs, total, nil
}
