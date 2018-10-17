package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/p-oka/graphql-demo/server/usecase"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTask": &graphql.Field{
			Type:    TaskType,
			Args:    graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: usecase.CreateTask,
		},
	},
})