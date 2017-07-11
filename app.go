package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/tokopedia/gosample/database"
	"github.com/tokopedia/gosample/hello"
	"github.com/tokopedia/gosample/product"
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

	database.InitDB()
	err := database.DBPool.MainDB.Ping()
	if err != nil {
		fmt.Println(err)
	}

	database.InitRedis()

	hwm := hello.NewHelloWorldModule()

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	http.HandleFunc("/getProduct", product.GetProductHandler)

	http.HandleFunc("/redis/set", database.RedisSetHandler)
	http.HandleFunc("/redis/get", database.RedisGetHandler)
	go logging.StatsLog()

	log.Fatal(grace.Serve(":9000", nil))
}
