package main

import (
	"ether_todo/pkg/injector"
	"ether_todo/pkg/glue"
	v1 "ether_todo/pkg/todo/controller/v1"
	"github.com/jasonlvhit/gocron"
	"github.com/kataras/i18n"
	"strconv"
)

func main() {
	_, err := i18n.New(i18n.Glob("./locales/*/*"), "en-US")
	if err != nil {
		panic(err)
	}

	e := glue.DefaultHttpServer()
	e = v1.Endpoints(e)
	startCron()
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(glue.ProvideServiceCfg().Port)))
}

func startCron() {
	err := gocron.Every(1).Day().Do(injector.TodoUsecase().CleanOutatedLists)
	if err != nil {
		panic(err)
	}
	gocron.Start()
}
