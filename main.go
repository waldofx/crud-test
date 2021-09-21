package main

import (
	"crud-test/config"
	"crud-test/routes"
)

func main() {
	config.InitDB()
	config.InitMigrate()
	e := routes.NewRoutes()
	e.Start(":8000")
}