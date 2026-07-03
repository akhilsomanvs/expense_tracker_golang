package handlers

import (
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request)  {}
