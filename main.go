package main

import (
	"go_api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", handlers.Greetings)
	e.GET("/foo", handlers.GetBars)
	e.GET("/foo/:uuid", handlers.GetBarsId)
	e.GET("/foo/sum", handlers.GetBarsSum)
	e.DELETE("/foo/:uuid", handlers.DeleteBarsId)
	e.POST("/foo", handlers.CreateNewBar)
	e.Logger.Fatal(e.Start("localhost:1323"))
}
