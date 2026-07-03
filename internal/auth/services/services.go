package services

import (
	"errors"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/entities"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/repositories"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Signup(req entities.SignUpRequest) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("Could not create Hash")
	}

	userModel := models.UserModel{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	err = s.repo.CreateUser(userModel)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) Login(req entities.LoginRequest) (string, error) {
	userModel, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid Email ID")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(userModel.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return "", errors.New("Invalid Credentials")
	}
	return "token", nil
}
