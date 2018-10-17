package gql

import (
	"github.com/graphql-go/graphql"
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