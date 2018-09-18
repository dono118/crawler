package main

import (
	"imooc.com/crawler/crawler-distributed/config"
	"imooc.com/crawler/crawler-distributed/rpcsupport"
	"imooc.com/crawler/crawler/engine"
	"imooc.com/crawler/crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1552811555",
		Type: "zhenai",
		Id:   "1552811555",
		Payload: model.Profile{
			Name:          "芜湖小啊妹",
			Gender:        "女",
			Age:           30,
			Height:        163,
			Weight:        0,
			Income:        "5001-8000元",
			Marriage:      "离异",
			Education:     "高中及以下",
			Occupation:    "销售专员",
			NativePlace:   "安徽芜湖",
			Constellation: "狮子座",
			House:         "和家人同住",
			Car:           "未购车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc,
		item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
