package server

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/p-oka/graphql-demo/server/gql"
)

func NewHandler() http.Handler {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    gql.QueryType,
		Mutation: gql.MutationType,
	})

	if err != nil {
		panic(err)
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: true,
	})
}
