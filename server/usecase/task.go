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

func CreateTask(p graphql.ResolveParams) (interface{}, error) {
	sess := p.Context.Value("db").(*dbr.Session)

	result, err := sess.InsertInto("tasks").Pair("name", p.Args["name"]).Exec()

	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	var task entity.Task
	_, err = sess.Select("*").From("tasks").Where("id = ?", lastInsertId).Load(&task)

	if err != nil {
		return nil, err
	}

	return task, nil
}