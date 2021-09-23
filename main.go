package main

import (
	"crud-test/config"
	ownMid "crud-test/middleware"
	"crud-test/routes"
)

func main() {
	config.InitDB()
	config.InitMigrate()
	e := routes.NewRoutes()
	ownMid.LogMiddlewareInit(e)
	e.Start(":8000")
}