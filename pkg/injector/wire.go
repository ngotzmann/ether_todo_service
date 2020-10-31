//+build wireinject

package injector

import (
	"ether_todo/pkg/controller/persistence"
	"ether_todo/pkg/todo"
	"github.com/google/wire"
)

func TodoUsecase() todo.IUsecase {
	wire.Build(persistence.NewTodoListRepo, todo.NewService, todo.NewUsecase)
	return &todo.Usecase{}
}
