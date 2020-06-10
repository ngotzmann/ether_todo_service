package main

import (
	v1 "ether_todo/pkg/controller/v1"
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

 	e := v1.EchoHandler("config/")
	e.Logger.Fatal(e.Start(c.Server.Address + ":" + c.Server.Port))
}
