package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/akhilsomanvs/expense_tracker/pkg/response"
)

func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(
						"Panic Recovered",
						slog.Any("error", err),
						slog.String("method", r.Method),
						slog.String("path", r.URL.Path),
					)
					response.InternalServerError[any](w)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
