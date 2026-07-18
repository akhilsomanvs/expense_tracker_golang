package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/configs"
	"github.com/akhilsomanvs/expense_tracker/internal/auth"
	"github.com/akhilsomanvs/expense_tracker/internal/core"
	"github.com/akhilsomanvs/expense_tracker/pkg/database"
	"github.com/akhilsomanvs/expense_tracker/pkg/jwt"
	"github.com/akhilsomanvs/expense_tracker/pkg/logger"
	"github.com/akhilsomanvs/expense_tracker/pkg/middlewares"
	"github.com/akhilsomanvs/expense_tracker/pkg/response"
	"github.com/go-chi/chi/v5"
)

func main() {

	config := configs.Load()
	jwtService := jwt.New(
		config.JWT.Secret,
		config.JWT.Issuer,
		config.JWT.TTL,
	)

	fmt.Println(config.Database.Host)

	fmt.Println(config.DatabaseURL())

	db, err := database.New(config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := chi.NewRouter()
	appLogger := logger.New()
	router.Use(
		middlewares.Recovery(appLogger),
		middlewares.RequestID,
		middlewares.Logging(appLogger),
	)

	//Module registry
	requiredModules := []core.AppModule{}

	modules := append(requiredModules, []core.AppModule{
		auth.NewModule(db.Pool(), jwtService),
	}...)
	for _, module := range modules {
		module.RegisterRoutes(router)
	}

	//Server
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("testing recovery middleware")
	})
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		response.OK(w, map[string]string{
			"message": "Hello",
		})
	})
	http.ListenAndServe(":8080", router)
}
