package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Code: 0, Message: "saved"})
}

func updateUser(c echo.Context) error {
	response := new(response)
	response.Code = 0
	response.Message = "updated"
	return c.JSON(http.StatusOK, response)
}

func deleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Code: 0, Message: "deleted"})
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}
