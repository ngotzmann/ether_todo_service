package main

import (
	v1 "ether_todo/pkg/controller/v1"
	"ether_todo/pkg/injector"
	"ether_todo/pkg/modules"
	"github.com/jasonlvhit/gocron"
	"github.com/kataras/i18n"
)

func main() {
	_, err := i18n.New(i18n.Glob("./locales/*/*"), "en-US")
	if err != nil {
		panic(err)
	}
	e := modules.DefaultHttpServer()
	e = v1.Endpoints(e)
	startCron()
	a := modules.ProvideServiceCfg().Port
	e.Logger.Fatal(e.Start(":" + a))
}

func startCron() {
	uc := injector.TodoUsecase()
	gocron.Every(1).Day().Do(uc.CleanOutatedLists)
	gocron.Start()
}
