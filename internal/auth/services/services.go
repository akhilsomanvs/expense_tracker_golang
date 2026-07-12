package services

import (
	"context"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/entities"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/repositories"
	"github.com/akhilsomanvs/expense_tracker/pkg/appErrors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository repositories.UserRepository
}

func New(repo repositories.UserRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Register(ctx context.Context, request entities.SignUpRequest) (*entities.SignUpResponse, error) {

	existingUser, err := s.repository.FindByEmail(
		ctx,
		request.Email,
	)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, appErrors.ErrEmailAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.UserModel{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hash),
	}

	err = s.repository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &entities.SignUpResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Message: "Registration successful",
	}, nil
}
