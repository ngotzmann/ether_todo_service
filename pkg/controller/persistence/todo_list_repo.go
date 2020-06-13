package persistence

import (
	"ether_todo/pkg/todo"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ngotzmann/gormmon"
	"github.com/ngotzmann/gorror"
	"github.com/sirupsen/logrus"
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

func (t *todoListRepo) DeleteOutdatedLists() {
	db, err := gormmon.GetGormDB()
	if err != nil {
		err = gorror.CreateError(gorror.DatabaseError, err.Error())
		logrus.Error(err)
	}
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '1 day'", todo.Day).Delete(&todo.List{})
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '30 day'", todo.Month).Delete(&todo.List{})
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '365 day'", todo.Year).Delete(&todo.List{})
}

func (t *todoListRepo) Migration() error {
	db, err := gormmon.GetGormDB()
	if err != nil {
		err = gorror.CreateError(gorror.DatabaseError, err.Error())
		return err
	}
	db.AutoMigrate(&todo.List{}, &todo.Task{})
	db.Model(&todo.Task{}).AddForeignKey("list_id", "lists(id)", "CASCADE", "CASCADE")
	return nil
}
