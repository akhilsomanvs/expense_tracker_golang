package repositories

import (
	"context"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserModel) error
	FindByEmail(ctx context.Context, email string) (*models.UserModel, error)
}
