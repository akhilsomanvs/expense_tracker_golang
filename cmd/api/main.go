package main

import (
	"context"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/internal/auth"
	"github.com/akhilsomanvs/expense_tracker/internal/core"
	"github.com/akhilsomanvs/expense_tracker/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {

	ctx := context.Background()
	storageModule := storage.NewModule(ctx)
	defer storageModule.Close()

	router := chi.NewRouter()

	requiredModules := []core.AppModule{
		storageModule,
	}

	modules := append(requiredModules, []core.AppModule{
		auth.NewModule(storageModule.ConnPool),
	}...)
	for _, module := range modules {
		module.RegisterRoutes(router)
	}
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", router)
}
