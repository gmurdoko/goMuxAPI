package main

import (
	"goMuxAPI/config"
	"goMuxAPI/main/master"
)

func main() {
	db := config.EnvConn()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
