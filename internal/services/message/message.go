package message

import (
	"context"
	"fmt"
	"time"

	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/3bd-dev/wallet-service/pkg/database"
	"github.com/3bd-dev/wallet-service/pkg/errs"
	"github.com/3bd-dev/wallet-service/pkg/logger"
	"github.com/3bd-dev/wallet-service/pkg/page"
	"github.com/google/uuid"
)

const (
	approvalsRequired = 2
)

type IMessageRepo interface {
	Create(ctx context.Context, message *models.Message) error
	QueryByID(ctx context.Context, id uuid.UUID) (*models.Message, error)
	Update(ctx context.Context, message *models.Message) error
	Query(ctx context.Context, filter QueryFilter, page page.Page) ([]models.Message, int64, error)
}

type IUserRepo interface {
	QueryByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type IMessageReviewRepo interface {
	Create(ctx context.Context, review *models.MessageReview) error
	CountApprovals(ctx context.Context, messageID uuid.UUID) (int, error)
	HasUserReviewed(ctx context.Context, messageID, userID uuid.UUID) (bool, error)
}

type Service struct {
	log               *logger.Logger
	messageRepo       IMessageRepo
	userRepo          IUserRepo
	messageReviewRepo IMessageReviewRepo
	transactor        database.ITransactor
}

func NewService(log *logger.Logger, mesgRepo IMessageRepo, usrRepo IUserRepo, msgReviewRepo IMessageReviewRepo, transactor database.ITransactor) *Service {
	return &Service{
		log:               log,
		messageRepo:       mesgRepo,
		userRepo:          usrRepo,
		messageReviewRepo: msgReviewRepo,
		transactor:        transactor,
	}
}

// Create creates a new message
func (s *Service) Create(ctx context.Context, r *NewMessage) (*models.Message, error) {
	if err := errs.Check(r); err != nil {
		return nil, err
	}

	usr, err := s.userRepo.QueryByID(ctx, r.SenderID)
	if err != nil {
		return nil, err
	}

	if usr.Role != models.RoleMaker {
		return nil, errs.New(errs.PermissionDenied, fmt.Errorf("user with ID %v is not allowed to create message", usr.ID))
	}

	message := &models.Message{
		ID:                uuid.New(),
		SenderID:          usr.ID,
		Content:           r.Content,
		Recipient:         r.Recipient,
		Status:            models.MessageStatusPending,
		ApprovalsRequired: approvalsRequired,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = s.messageRepo.Create(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *Service) Review(ctx context.Context, messageID uuid.UUID, r *NewMessageReview) (*models.MessageReview, error) {
	var err error
	if err = errs.Check(r); err != nil {
		return nil, err
	}

	tx := s.transactor.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dbctx := context.WithValue(ctx, database.ContextKeyDBTx, tx)

	reviewer, err := s.userRepo.QueryByID(ctx, r.ReviewerID)
	if err != nil {
		return nil, err
	}

	if reviewer.Role != models.RoleChecker {
		return nil, errs.New(errs.PermissionDenied, fmt.Errorf("user with ID %v is not allowed to review", reviewer.ID))
	}

	message, err := s.messageRepo.QueryByID(dbctx, messageID)
	if err != nil {
		return nil, err
	}

	if message.Status != models.MessageStatusPending {
		return nil, errs.New(errs.InvalidArgument, fmt.Errorf("message is already %s", message.Status))
	}

	if reviewer.ID == message.SenderID {
		return nil, errs.New(errs.PermissionDenied, fmt.Errorf("user with ID %v is not allowed to review their own message", reviewer.ID))
	}

	// Check if the reviewer already reviewed the message
	hasReviewed, err := s.messageReviewRepo.HasUserReviewed(dbctx, messageID, r.ReviewerID)
	if err != nil {
		return nil, err
	}
	if hasReviewed {
		return nil, errs.New(errs.InvalidArgument, fmt.Errorf("user %v has already reviewed this message", r.ReviewerID))
	}

	review := &models.MessageReview{
		ID:         uuid.New(),
		MessageID:  messageID,
		ReviewerID: r.ReviewerID,
		Action:     r.Action,
		CreatedAt:  time.Now(),
	}

	err = s.messageReviewRepo.Create(dbctx, review)
	if err != nil {
		return nil, err
	}

	// Check review action
	if r.Action == models.MessageReviewActionApproved {
		// Count approvals
		approvalCount, err := s.messageReviewRepo.CountApprovals(dbctx, messageID)
		if err != nil {
			return nil, err
		}

		// If approvals meet the required number, approve the message
		if approvalCount >= message.ApprovalsRequired {
			message.Status = models.MessageStatusApproved
		}
	} else {
		message.Status = models.MessageStatusRejected
	}

	// Update message status only if it changed
	if message.Status != models.MessageStatusPending {
		message.UpdatedAt = time.Now()
		err = s.messageRepo.Update(dbctx, message)
		if err != nil {
			return nil, err
		}
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return nil, commitErr
	}

	return review, nil
}

// Query to list messages
func (s *Service) Query(ctx context.Context, filter QueryFilter, page page.Page) ([]models.Message, int64, error) {
	msgs, total, err := s.messageRepo.Query(ctx, filter, page)
	if err != nil {
		return nil, 0, errs.New(errs.Internal, err)
	}

	return msgs, total, nil
}
