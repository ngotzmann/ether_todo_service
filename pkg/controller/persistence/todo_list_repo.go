package persistence

import (
	"ether_todo/pkg/controller/persistence/gormmon"
	"ether_todo/pkg/todo"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ngotzmann/gorror"
)


type todoListRepo struct {
}

func NewTodoListRepo() *todoListRepo {
	return &todoListRepo{}
}

func (t *todoListRepo) FindListByName(name string) (*todo.List, error) {
	db, err := gormmon.GetGormDB()
	if err != nil {
		err = gorror.CreateError(gorror.DatabaseError, err.Error())
		return nil, err
	}

	l := &todo.List{}
	db.Where("name = ?", name).Preload("Tasks").Find(&l)

	if  l.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, nil
	} else {
		return l, nil
	}
}

func (t *todoListRepo) SaveList(l *todo.List) (*todo.List, error) {
	db, err := gormmon.GetGormDB()
	if err != nil {
		err = gorror.CreateError(gorror.DatabaseError, err.Error())
		return nil, err
	}

	db.Save(&l)
	return l, nil
}

func (t *todoListRepo) DeleteListByName(l *todo.List) error {
	db, err := gormmon.GetGormDB()
	if err != nil {
		err = gorror.CreateError(gorror.DatabaseError, err.Error())
		return err
	}
	db.Unscoped().Where("name = ?", l.Name).Delete(&l)
	return nil
}

/*
ALTER TABLE tasks ADD CONSTRAINT list_fkey FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE;
db.AutoMigrate(&todo.List{}, &todo.Task{})
	db.Create(&todo.List{
		ID:        uuid.MustParse("696f7dcd-ce91-4bea-ac87-62dfe33b8329"),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      "ketchup",
	//	Tasks:     nil,
		LiveTime:  "keep",
	})
*/