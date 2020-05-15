package main

import (
	v1 "boilerplate_service/pkg/controller/v1"

	"github.com/ngotzmann/gommon"
	"github.com/ngotzmann/gorror"
)

func main() {
	gorror.Init("config/")
	c := gommon.NewConfig("config/")
	gommon.InitLogrus(c.Logging.Level, c.Logging.File, c.Logging.TimestampFormat)

	e := v1.EchoHandler("./")
	e.Logger.Fatal(e.Start(c.Server.Address + ":" + c.Server.Port))
}
