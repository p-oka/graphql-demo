package usecase

import (
	"github.com/graphql-go/graphql"
	"github.com/p-oka/graphql-demo/server/entity"
)

func GetTasks(p graphql.ResolveParams) (interface{}, error) {
	return []entity.Task{
		entity.Task{ID: 1, Name: "osomatsu"},
		entity.Task{ID: 2, Name: "karamatsu"},
		entity.Task{ID: 3, Name: "choromatsu"},
		entity.Task{ID: 4, Name: "ichimatsu"},
	}, nil
}