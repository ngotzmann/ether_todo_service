package main

import (
	v1 "ether_todo/pkg/controller/v1"
	"ether_todo/pkg/injector"
	"ether_todo/pkg/modules"
	"github.com/jasonlvhit/gocron"
	"github.com/ngotzmann/gorror"
)

func main() {
	gorror.Init(modules.DefaultConfig().GorrorFilePath)
	e := modules.DefaultHttpServer()
	e = v1.Endpoints(e)
	startCron()
	e.Logger.Fatal(e.Start(":"+modules.DefaultConfig().Port))
}

func startCron() {
	uc := injector.TodoUsecase()
 	gocron.Every(1).Day().Do(uc.CleanOutatedLists)
 	gocron.Start()
 }
