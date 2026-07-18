package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/internal/auth/entities"
	"github.com/akhilsomanvs/expense_tracker/pkg/appErrors"
	"github.com/akhilsomanvs/expense_tracker/pkg/response"
	"github.com/akhilsomanvs/expense_tracker/pkg/validator"
)

type Handler struct {
	service AuthService
}

func NewHandler(service AuthService) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var request entities.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.Validate(request); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	registerResponse, err := h.service.Register(r.Context(), request)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrEmailAlreadyExists):
			response.Conflict[string](w, err.Error())
		default:
			response.InternalServerError[any](w)
		}
		return
	}

	response.Created(w, registerResponse)
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var request entities.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.Validate(request); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	loginResponse, err := h.service.Login(r.Context(), request)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidCredentials):
			response.Unauthorized[any](w, "Invalid email or password")

		default:
			response.InternalServerError[any](w)
		}

		return
	}
	response.OK(
		w,
		loginResponse,
	)
}
