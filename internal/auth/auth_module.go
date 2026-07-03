package auth

import (
	"github.com/akhilsomanvs/expense_tracker/internal/auth/handlers"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/repositories"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/routes"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/services"
	"github.com/go-chi/chi/v5"
)

type AuthModule struct {
	ModuleName string
	routes.AuthRoutes
	handler *handlers.Handler
}

func NewModule() *AuthModule {
	repo := repositories.NewRepository()
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)
	return &AuthModule{
		ModuleName: "auth",
		AuthRoutes: routes.AuthRoutes{},
		handler:    handler,
	}
}

func (m *AuthModule) Name() string {
	return m.ModuleName
}

func (m *AuthModule) RegisterRoutes(router chi.Router) {
	m.AuthRoutes.Routes()
	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", m.handler.Signup)
		r.Post("/login", m.handler.Login)
	})
}
