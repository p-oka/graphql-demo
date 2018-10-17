package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/p-oka/graphql-demo/server/usecase"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"tasks": &graphql.Field{
			Type:    graphql.NewList(TaskType),
			Resolve: usecase.GetTasks,
		},
	},
})