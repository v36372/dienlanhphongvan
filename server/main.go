package main

import (
	"dienlanhphongvan/app/handler"
	"dienlanhphongvan/cmd"
	"dienlanhphongvan/config"
	"dienlanhphongvan/infra"
	"dienlanhphongvan/utilities/ulog"
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	cmd.Execute()
	fmt.Println("about to get")
	conf := config.Get()

	setupInfra(conf)
	defer infra.ClosePostgreSql()

	ginEngine := handler.InitEngine(&conf)
	address := fmt.Sprintf("%s:%d", config.Get().App.Host, config.Get().App.Port)
	server := http.Server{
		Addr:    address,
		Handler: ginEngine,
	}

	if err := gracehttp.Serve(&server); err != nil {
		panic(err)
	}
}

func setupInfra(conf config.Config) {
	// Init logger
	ulog.InitDefaultLogger(conf.Log.Dir, conf.Log.LevelDebug)

	// Postgresql
	infra.InitPostgreSQL()
}
