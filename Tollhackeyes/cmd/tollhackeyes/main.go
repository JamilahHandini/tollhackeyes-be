package main

import (
	"github.com/jamilahhandini/tollhackeyes/config"
	"github.com/jamilahhandini/tollhackeyes/server"
	dbf "github.com/jamilahhandini/tollhackeyes/dbf"
)

func main() {
	//* ====================== Config ======================
	conf := config.InitConfig("local")

	//* ====================== Database ======================
	fb := dbf.NewConnection(conf)

	//* ====================== Running Server ======================
	server.Run(conf, fb)
}
