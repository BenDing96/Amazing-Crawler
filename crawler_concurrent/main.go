package main

import (
	"learngo/crawler_official/config"
	"learngo/crawler_official/engine"
	"learngo/crawler_official/persist"
	"learngo/crawler_official/scheduler"
	"learngo/crawler_official/xcar/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://newcar.xcar.com.cn/car/",
		Parser: engine.NewFuncParser(
			parser.ParseCarList,
			config.ParseCarList),
	})
}
