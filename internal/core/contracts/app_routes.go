package contracts

import "github.com/go-chi/chi/v5"

type AppRoutes interface {
	RegisterRoutes(router chi.Router)
}
