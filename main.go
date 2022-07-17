package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Debug
	e.Debug = true

	// Middleware
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

	// Handling Request
	e.POST("/users", users)

	// Static Content
	e.Static("/static", "static")

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "admin login")
	})

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.Logger.Fatal(e.Start(":1323"))
}
