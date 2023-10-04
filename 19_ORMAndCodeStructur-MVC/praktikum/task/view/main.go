package main

import (
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/config"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/route"
)

func main() {
	// init DB
	config.InitDB()

	// create a new echo instance
	e := route.NewRoute()


	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
