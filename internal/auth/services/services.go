package services

import (
	"context"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/entities"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/repositories"
	"github.com/akhilsomanvs/expense_tracker/pkg/appErrors"
	"github.com/akhilsomanvs/expense_tracker/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository repositories.UserRepository
	jwtService *jwt.Manager
}

func New(repo repositories.UserRepository, jwtService *jwt.Manager) *Service {
	return &Service{
		repository: repo,
		jwtService: jwtService,
	}
}

func (s *Service) Register(ctx context.Context, request entities.SignUpRequest) (*entities.SignUpResponse, error) {

	existingUser, err := s.repository.GetUserByEmail(
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

func (s *Service) Login(ctx context.Context, request entities.LoginRequest) (*entities.LoginResponse, error) {
	user, err := s.repository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, appErrors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		return nil, appErrors.ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(
		&jwt.TokenPayload{
			UserID: user.ID,
			Email:  user.Email,
		},
	)

	if err != nil {
		return nil, err
	}

	return &entities.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(s.jwtService.TTL().Seconds()),
	}, nil
}
