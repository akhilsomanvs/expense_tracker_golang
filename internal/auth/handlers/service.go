package handlers

import (
	"context"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/entities"
)

type AuthService interface {
	Register(ctx context.Context, request entities.SignUpRequest) (*entities.SignUpResponse, error)
}
