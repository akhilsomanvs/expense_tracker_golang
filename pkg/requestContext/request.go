package requestcontext

import "context"

func RequestID(ctx context.Context) string {

	requestID, ok := ctx.Value(RequestIDKey).(string)
	if !ok {
		return ""
	}

	return requestID
}

func UserID(ctx context.Context) string {

	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return ""
	}

	return userID
}

func UserEmail(ctx context.Context) string {

	email, ok := ctx.Value(UserEmailKey).(string)
	if !ok {
		return ""
	}

	return email
}
