package main

import (
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/internal/auth"
	"github.com/akhilsomanvs/expense_tracker/internal/core"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	modules := []core.AppModule{
		auth.NewModule(),
	}
	for _, module := range modules {
		module.RegisterRoutes(router)
	}
	http.ListenAndServe(":8080", router)
}
