package main

import (
	"github.com/mychengel/depd_go_echo/db"
	"github.com/mychengel/depd_go_echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":9999"))
}
