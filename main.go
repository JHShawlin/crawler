package main

import (
	"crawler/engine"
	"crawler/niuke/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 10,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
