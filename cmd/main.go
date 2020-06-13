package main

import (
	"ether_todo/pkg/controller/persistence"
	v1 "ether_todo/pkg/controller/v1"
	"ether_todo/pkg/todo"
	"github.com/jasonlvhit/gocron"
	"github.com/ngotzmann/gormmon"

	"github.com/ngotzmann/gommon"
	"github.com/ngotzmann/gorror"
)

func main() {
	gorror.Init("config/")
	c := gommon.NewConfig("config/")
	gommon.InitLogrus(c.Logging.Level, c.Logging.File, c.Logging.TimestampFormat)
	gormmon.InitGormDB(gormmon.GormConfig{
		Host:               c.Database.Address,
		Port:               c.Database.Port,
		DBName:             c.Database.Database,
		Username:           c.Database.User,
		Password:           c.Database.Password,
		MaxIdleConnections: c.Database.MaxIdleConnections,
		ShouldLog:          c.Database.Logging,
	})

	startGocron()

	e := v1.EchoHandler("config/")
	e.Logger.Fatal(e.Start(c.Server.Address + ":" + c.Server.Port))
}

func startGocron() {
	repo := persistence.NewTodoListRepo()
	uc := todo.NewUsecase(repo, todo.NewService(repo))
	gocron.Every(1).Day().Do(uc.CleanOutatedLists)
	gocron.Start()
}