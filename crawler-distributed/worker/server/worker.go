package main

import (
	"flag"
	"fmt"
	"imooc.com/crawler/crawler-distributed/rpcsupport"
	"imooc.com/crawler/crawler-distributed/worker"
	"log"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
