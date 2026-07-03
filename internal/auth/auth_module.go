package auth

import (
	"github.com/akhilsomanvs/expense_tracker/internal/auth/routes"
	"github.com/go-chi/chi/v5"
)

type AuthModule struct {
	ModuleName string
	routes.AuthRoutes
}

func NewModule() *AuthModule {
	return &AuthModule{
		ModuleName: "auth",
		AuthRoutes: routes.AuthRoutes{},
	}
}

func (m *AuthModule) Name() string {
	return m.ModuleName
}

func (m *AuthModule) RegisterRoutes(router chi.Router) {
	m.AuthRoutes.Routes()
}
