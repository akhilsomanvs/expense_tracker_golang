package middlewares

import (
	"context"
	"net/http"

	requestcontext "github.com/akhilsomanvs/expense_tracker/pkg/requestContext"
	"github.com/google/uuid"
)

const HeaderRequestID = "X-Request-ID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		ctx := context.WithValue(
			r.Context(),
			requestcontext.RequestIDKey,
			requestID,
		)

		w.Header().Set(HeaderRequestID, requestID)
		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
