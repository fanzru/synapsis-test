package contextutil

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	UserID   contextKey = "user_id"
	Username contextKey = "username"
)

func OverrideContext(echoCtx echo.Context) context.Context {
	ctx := echoCtx.Request().Context()

	ctx = context.WithValue(ctx, UserID, echoCtx.Get(string(UserID)))
	ctx = context.WithValue(ctx, Username, echoCtx.Get(string(Username)))

	return ctx
}

func GetUserID(ctx context.Context) int64 {
	if num, ok := ctx.Value(UserID).(int64); ok {
		return num
	}
	return 0
}
