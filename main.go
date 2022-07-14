package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/user/:id", getUser)
	e.POST("/user", saveUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	e.GET("/show", show)

	e.Logger.Fatal(e.Start(":1323"))
}
