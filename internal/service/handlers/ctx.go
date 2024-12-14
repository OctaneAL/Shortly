package handlers

import (
	"context"
	"net/http"

	"github.com/OctaneAL/Shortly/internal/db"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int
type contextKey string

const (
	logCtxKey ctxKey     = iota
	dbKey     contextKey = "db"
)

func CtxDB(ctx context.Context, database *db.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbKey, database)
	}
}

func DB(ctx context.Context) *db.DB {
	value := ctx.Value(dbKey)
	if value == nil {
		panic("attempt to retrieve DB from context, but it is not set")
	}
	return value.(*db.DB)
}

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
