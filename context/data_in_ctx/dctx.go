package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("fero", "fero123")
}

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(ctx context.Context) string {
	return ctx.Value(ctxUserID).(string)
}

func AuthToken(ctx context.Context) string {
	return ctx.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	// note that key and values are stored as interface{}
	// recommendations to overcome this:
	// 1. create a custom key type in your package
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response %v (%v)\n",
		// ctx.Value("userID"),
		// ctx.Value("authToken"),
		UserID(ctx), // used for type safety - similar as getter in Java
		AuthToken(ctx),
	)
}
