package main

import (
	"golang-daily/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/diaries", handlers.GetDiaries)
	e.POST("/diaries/new", handlers.CreateDiary)
	e.GET("/diaries/:id", handlers.GetDiary)
	e.PUT("/diaries/:id", handlers.UpdateDiary)
	e.DELETE("/diaries/:id", handlers.DeleteDiary)
	e.Logger.Fatal(e.Start(":8080"))
}
