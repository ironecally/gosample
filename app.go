package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/tokopedia/gosample/db"
	"github.com/tokopedia/gosample/hello"
	"github.com/tokopedia/gosample/product"
	"github.com/tokopedia/gosample/redis"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	// gops helps us get stack trace if something wrong/slow in production
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}

	db.InitDB()
	redis.InitRedis()

	hwm := hello.NewHelloWorldModule()

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	http.HandleFunc("/product", product.GetProductHandler)

	http.HandleFunc("/redis/set", redis.SetRedisHandler)
	http.HandleFunc("/redis/get", redis.GetRedisHandler)

	go logging.StatsLog()

	log.Fatal(grace.Serve(":9000", nil))
}
