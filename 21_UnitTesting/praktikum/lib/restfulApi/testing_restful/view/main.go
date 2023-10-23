package main

import (
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/route"
)

func main() {
	// init DB
	config.InitDB()

	// create a new echo instance
	e := route.NewRoute()


	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
