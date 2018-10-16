package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/p-oka/graphql-demo/server/usecase"
)

var TaskType = graphql.NewObject(graphql.ObjectConfig{
		Name: "task",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
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