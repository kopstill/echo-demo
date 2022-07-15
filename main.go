package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Hello, World!
	e.GET("/", hello)

	// Routing
	// Path Parameters
	e.GET("/user/:id", getUser)
	e.POST("/user", saveUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	// Query Parameters
	e.GET("/show", show)

	//Form application/x-www-form-urlencoded
	e.POST("/save", save)

	// Form multipart/form-data
	e.POST("/avatar", avatar)

	e.Logger.Fatal(e.Start(":1323"))
}
