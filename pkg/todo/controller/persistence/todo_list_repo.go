package persistence

import (
	"errors"
	"ether_todo/pkg/glue"
	"ether_todo/pkg/todo"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/i18n"
	"github.com/labstack/gommon/log"
)

type todoListRepo struct {
}

func NewTodoListRepo() todo.IRepository {
	return &todoListRepo{}
}

func (t *todoListRepo) FindListByName(name string) (*todo.List, error) {
	db, err := glue.DefaultGorm()
	if err != nil {
		err := errors.New(i18n.Tr("en-US", "DatabaseError"))
		return nil, err
	}

	l := &todo.List{}
	db.Where("name = ?", name).Preload("Tasks").Find(&l)

	if l.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, nil
	} else {
		return l, nil
	}
}

func (t *todoListRepo) SaveList(l *todo.List) (*todo.List, error) {
	db, err := glue.DefaultGorm()
	if err != nil {
		return nil, err
	}
	if l.ID.String() == "00000000-0000-0000-0000-000000000000"  {
		l.ID = uuid.New()
	}

	db.Save(&l)
	return l, nil
}

func (t *todoListRepo) DeleteListByName(l *todo.List) error {
	db, err := glue.DefaultGorm()
	if err != nil {
		return err
	}
	db.Unscoped().Where("name = ?", l.Name).Delete(&l)
	return nil
}

func (t *todoListRepo) DeleteOutdatedLists() {
	db, err := glue.DefaultGorm()
	if err != nil {
		log.Error(err)
	}
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '1 day'", todo.Day).Delete(&todo.List{})
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '30 day'", todo.Month).Delete(&todo.List{})
	db.Unscoped().Where("live_time = ? AND updated_at < CURRENT_TIMESTAMP - INTERVAL '365 day'", todo.Year).Delete(&todo.List{})
}

func (t *todoListRepo) Migration() error {
	db, err := glue.DefaultGorm()
	if err != nil {
		return err
	}
	db.AutoMigrate(&todo.List{}, &todo.Task{})
	db.Model(&todo.Task{}).AddForeignKey("list_id", "lists(id)", "CASCADE", "CASCADE")
	return nil
}
