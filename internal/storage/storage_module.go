package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akhilsomanvs/expense_tracker/internal/storage/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageModule struct {
	ModuleName string
	handler    *handlers.Handler
	ConnPool   *pgxpool.Pool
}

func NewModule(ctx context.Context) *StorageModule {
	pool, err := newPool(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to DB : %v", err)
	}
	handler := &handlers.Handler{
		ConnPool: pool,
	}
	return &StorageModule{
		ModuleName: "storage",
		handler:    handler,
		ConnPool:   pool,
	}
}

func (sm *StorageModule) Close() {

}

func newPool(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 20
	config.MinConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func (m *StorageModule) Name() string {
	return m.ModuleName
}

func (m *StorageModule) RegisterRoutes(router chi.Router) {
	router.Route("/storage", func(r chi.Router) {
		r.Get("/healthCheck", m.handler.DBHealthCheck)
	})
}
