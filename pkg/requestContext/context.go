package requestcontext

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserIDKey    contextKey = "user_id"
	UserEmailKey contextKey = "user_email"
)
