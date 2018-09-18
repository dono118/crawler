package main

import (
	"flag"
	"imooc.com/crawler/crawler-distributed/config"
	itemSaver "imooc.com/crawler/crawler-distributed/persist/client"
	"imooc.com/crawler/crawler-distributed/rpcsupport"
	worker "imooc.com/crawler/crawler-distributed/worker/client"
	"imooc.com/crawler/crawler/engine"
	"imooc.com/crawler/crawler/scheduler"
	"imooc.com/crawler/crawler/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "","itemsaver host")
	workerHosts = flag.String("worker_hosts", "",
		"worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList,
			config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
