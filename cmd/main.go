package main

import (
	v1 "ether_todo/pkg/adapter/v1"
	"ether_todo/pkg/glue"
	"ether_todo/pkg/injector"
	"github.com/jasonlvhit/gocron"
	"github.com/kataras/i18n"
	"strconv"
)

func main() {
	_, err := i18n.New(i18n.Glob("./locales/*/*"), "en-US")
	if err != nil {
		panic(err)
	}

	err = injector.TodoService().Migration()
	if err != nil {
		panic(err)
	}

	e := glue.DefaultHttpServer()
	e = v1.Endpoints(e)
	startCron()
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(glue.ProvideServiceCfg().Port)))
}

func startCron() {
	err := gocron.Every(1).Day().Do(injector.TodoService().CleanOutatedLists)
	if err != nil {
		panic(err)
	}
	gocron.Start()
}
