//+build wireinject

package injector

import (
	"ether_todo/pkg/todo"
	"ether_todo/pkg/todo/controller/persistence"
	"github.com/google/wire"
)

func TodoUsecase() todo.IUsecase {
	wire.Build(persistence.NewTodoListRepo, todo.NewService, todo.NewUsecase)
	return &todo.Usecase{}
}
