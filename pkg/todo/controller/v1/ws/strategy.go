package ws

import (
	"fmt"
	"reflect"
)

//Util BEGIN
//https://play.golang.org/
var typeRegistry = make(map[string]reflect.Type)

func RegisterType(i interface{}) {
	t := reflect.TypeOf(i).Elem()
	typeRegistry[t.PkgPath() + "." + t.Name()] = t
}

func MakeInstance(name string) interface{} {
	return reflect.New(typeRegistry[name]).Elem().Interface()
}
//Util END


type WebsocketStrategy interface {
	RegisterType()
	Save()
}

type TodoTaskSave struct {}

func (t *TodoTaskSave) RegisterType() {
	RegisterType((*TodoTaskSave)(nil))
}

func (t *TodoTaskSave) Save(m Request) {
	fmt.Println("Save ToDoTask")
}