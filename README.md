# crawler
A distributed crawler based on Golang.

---


dependences:

```
// 1. docker 18.06.1-ce
// 2. 安装elasticsearch
docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2
// 3. 安装elastic client:
go get -v gopkg.in/olivere/elastic.v5
```

Run:

```
docker ps

// 若有(elasticsearch CONTAINER ID)
docker kill id

docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2

并发版爬虫
go run crawler/main.go

分布式爬虫
go run crawler-distributed/persist/itemsaver.go --port=1234
go run crawler-distributed/worker/server/worker.go --port=9000
go run crawler-distributed/worker/server/worker.go --port=9001
go run crawler-distributed/worker/server/worker.go --port=9002
go run crawler-distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001,:9002"


