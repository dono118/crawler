package main

import (
	"fmt"
	"imooc.com/crawler/crawler-distributed/config"
	"imooc.com/crawler/crawler-distributed/rpcsupport"
	"imooc.com/crawler/crawler-distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1552811555",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "芜湖小啊妹",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
