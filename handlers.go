package main

import (
	"io"
	"net/http"
	"os"
	"strings"

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

func avatar(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// dst, err := os.Create("/Users/kopever/Develop/temp/echo-demo/" + avatar.Filename)
	dst, err := os.Create("/Users/kopever/Develop/temp/echo-demo/" + name + "." + strings.Split(avatar.Filename, ".")[len(strings.Split(avatar.Filename, "."))-1])
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
