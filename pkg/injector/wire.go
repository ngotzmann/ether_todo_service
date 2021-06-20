//+build wireinject

package injector

import (
	"ether_todo/pkg/adapter/persistence"
	"ether_todo/pkg/todo"
	"github.com/google/wire"
)

func TodoService() todo.IService {
	wire.Build(persistence.NewTodoListRepo, todo.NewService)
	return &todo.Service{}
}
