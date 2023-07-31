package main

import (

	"github.com/jamilahhandini/tollhackeyes/config"
	dbf "github.com/jamilahhandini/tollhackeyes/dbf"
	"github.com/jamilahhandini/tollhackeyes/server"
)

func main() {
	//* ====================== Config ======================
	conf := config.InitConfig("local")

	//* ====================== Database ======================
	fb := dbf.NewConnection(conf)

	//* ====================== Running Server ======================
	server.Run(conf, fb)
}
