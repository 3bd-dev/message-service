package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/3bd-dev/wallet-service/internal/models"
	"github.com/3bd-dev/wallet-service/pkg/database"
	"github.com/3bd-dev/wallet-service/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db database.IDatabase
}

// NewUserRepo creates a new instance of UserRepo.
func NewUserRepo(db database.IDatabase) *UserRepo {
	return &UserRepo{db: db}
}

// QueryByID get User from database
func (r *UserRepo) QueryByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var usr models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&usr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.NotFound, fmt.Errorf("user with ID %s not found", id))
		}
		return nil, err
	}

	return &usr, nil
}
