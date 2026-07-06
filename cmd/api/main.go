package main

import (
	"context"
	"log"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/internal/auth"
	"github.com/akhilsomanvs/expense_tracker/internal/core"
	"github.com/akhilsomanvs/expense_tracker/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {

	ctx := context.Background()
	pool, err := storage.NewPool(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to DB : %v", err)
	}
	defer pool.Close()

	log.Println("Connected to DB")

	router := chi.NewRouter()

	requiredModules := []core.AppModule{
		storage.NewModule(pool),
	}

	modules := append(requiredModules, []core.AppModule{
		auth.NewModule(pool),
	}...)
	for _, module := range modules {
		module.RegisterRoutes(router)
	}
	http.ListenAndServe(":8080", router)
}
