package main

import (
	"condog/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.GetEndPointHandler)
	e.Logger.Fatal(e.Start(":80"))
}
