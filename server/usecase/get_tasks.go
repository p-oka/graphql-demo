package usecase

import (
	"github.com/graphql-go/graphql"
	"github.com/p-oka/graphql-demo/server/entity"
	"github.com/gocraft/dbr"
)

func GetTasks(p graphql.ResolveParams) (interface{}, error) {
	sess := p.Context.Value("db").(*dbr.Session)

	var tasks []entity.Task

	_, err := sess.Select("*").From("tasks").Load(&tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}