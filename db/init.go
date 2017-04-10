package db

import (
	"fmt"
	"log"

	logging "gopkg.in/tokopedia/logging.v1"

	_ "github.com/lib/pq"
	"github.com/tokopedia/sqlt"
)

var DBPools dbPool

type dbPool struct {
	DB1 *sqlt.DB
}

type Config struct {
	DB struct {
		DSN string
	}
}

func InitDB() {
	var dbCfg Config
	ok := logging.ReadModuleConfig(&dbCfg, "/etc/gosample", "db") || logging.ReadModuleConfig(&dbCfg, "files/etc/gosample", "db")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	fmt.Println("init db connection")

	dbConn, err := sqlt.Open("postgres", dbCfg.DB.DSN)
	if err != nil {
		log.Fatalln("failed to init db connection", err.Error())
	}

	err = dbConn.Ping()
	if err != nil {
		log.Println("db is not reachable", err.Error())
	}

	DBPools = dbPool{
		DB1: dbConn,
	}
}
