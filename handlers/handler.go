package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"golang-daily/models"
)

func GetDiaries(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func CreateDiary(c echo.Context) error {
	diary := new(models.Diary)
	if err := c.Request().ParseForm(); err != nil {
		fmt.Println("error")
	}
	for key, value := range c.Request().Form {
		fmt.Printf("%v : %v\n", key, value)
	}
	if err := c.Bind(diary); err != nil {
		return err
	}
	fmt.Println(diary)
	models.Db.Create(diary)
	return c.JSON(http.StatusCreated, diary)
}

func GetDiary(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}

func UpdateDiary(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}

func DeleteDiary(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}
