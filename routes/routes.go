package routes

import (
	"crud-test/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo{
	e := echo.New()
	e.GET("/user/:id",controllers.GetUserController)
	e.PUT("/user/:id",controllers.UpdateUserController)
	e.DELETE("/user/:id",controllers.DeleteUserController)
	e.GET("/user",controllers.GetAllUserController)
	e.POST("/user",controllers.AddUserController)
	return e
}