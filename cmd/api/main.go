package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/configs"
	"github.com/akhilsomanvs/expense_tracker/internal/auth"
	"github.com/akhilsomanvs/expense_tracker/internal/core"
	"github.com/akhilsomanvs/expense_tracker/pkg/database"
	"github.com/akhilsomanvs/expense_tracker/pkg/logger"
	"github.com/akhilsomanvs/expense_tracker/pkg/middlewares"
	"github.com/go-chi/chi/v5"
)

func main() {

	config := configs.Load()

	fmt.Println(config.Database.Host)

	fmt.Println(config.DatabaseURL())

	db, err := database.New(config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := chi.NewRouter()
	appLogger := logger.New()
	router.Use(middlewares.Recovery(appLogger))

	//Module registry
	requiredModules := []core.AppModule{}

	modules := append(requiredModules, []core.AppModule{
		auth.NewModule(db.Pool()),
	}...)
	for _, module := range modules {
		module.RegisterRoutes(router)
	}

	//Server
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", router)
}
