package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	mwinternal "github.com/akhilsomanvs/expense_tracker/pkg/middlewares/internal"
	requestcontext "github.com/akhilsomanvs/expense_tracker/pkg/requestContext"
)

func Logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := mwinternal.NewResponseWriter(w)

			next.ServeHTTP(rw, r)

			logger.Info(
				"request completed",
				slog.String("request_id", requestcontext.RequestID(r.Context())),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rw.StatusCode()),
				slog.Int("response_size", rw.Size()),
				slog.Duration("duration", time.Since(start)),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			)
		})
	}
}
