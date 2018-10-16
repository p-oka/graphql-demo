package server

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"net/http"
	"os"
)

func WithConnection(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := NewConnection()
		ctx := context.WithValue(r.Context(), "db", sess.NewSession(nil))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewConnection() *dbr.Connection {
	conn, err := dbr.Open("mysql", os.Getenv("dsn"), nil)

	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(5)

	return conn
}
