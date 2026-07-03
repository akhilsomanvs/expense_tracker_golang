package core

import "github.com/akhilsomanvs/expense_tracker/internal/core/contracts"

type AppModule interface {
	Name() string
	contracts.AppRoutes
}
