package main

import (
	"catalog/src/api"
	"catalog/src/config"
	"catalog/src/db/postgres"
	"catalog/src/logger"
	"catalog/src/server"
	"fmt"
)

func main() {
	conf, err := config.NewConfigFromEnvVars()
	if err != nil {
		panic(err)
	}
	log, err := logger.NewLogger(&conf.Logger)
	if err != nil {
		panic(err)
	}
	dbConnections, err := postgres.NewDBConnections(&conf.Postgres)
	if err != nil {
		panic(err)
	}
	fmt.Print(dbConnections)
	//storage := api.NewStorage(dbConnections)
	//controller := api.NewController(storage)
	app := api.NewApp(nil)
	router := app.NewRouter()
	cacheClient := server.NewRedisCache(conf)
	srv := server.NewServer(&conf.Server, log, cacheClient)
	srv.Run(router)
}
