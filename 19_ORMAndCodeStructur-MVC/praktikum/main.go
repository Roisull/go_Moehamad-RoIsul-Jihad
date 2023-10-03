package main

import (
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/config"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/route"
)

func main() {
	// memanggil initDB
	config.InitDB()

	e := route.New()
	e.Start(":8000")
}
