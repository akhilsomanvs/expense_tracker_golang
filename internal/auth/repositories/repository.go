package repositories

import (
	"errors"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/models"
)

type Repository struct {
	userModel models.UserModel
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUser(user models.UserModel) error {
	r.userModel = user
	return nil
}
func (r *Repository) GetUserByEmail(email string) (*models.UserModel, error) {
	if r.userModel.Email == email {
		return &r.userModel, nil
	}
	return &models.UserModel{}, errors.New("Could not find user")
}
