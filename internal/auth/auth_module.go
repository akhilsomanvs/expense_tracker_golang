package auth

import (
	"github.com/akhilsomanvs/expense_tracker/internal/auth/handlers"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/repositories"
	"github.com/akhilsomanvs/expense_tracker/internal/auth/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthModule struct {
	ModuleName string
	handler    *handlers.Handler
}

func NewModule(pool *pgxpool.Pool) *AuthModule {
	repo := repositories.NewPostgressRepository(pool)
	service := services.New(repo)
	handler := handlers.NewHandler(service)
	return &AuthModule{
		ModuleName: "auth",
		handler:    handler,
	}
}

func (m *AuthModule) Name() string {
	return m.ModuleName
}

func (m *AuthModule) RegisterRoutes(router chi.Router) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", m.handler.Signup)
		r.Post("/login", m.handler.Login)
	})
}
