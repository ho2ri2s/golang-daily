package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"golang-daily/models"
)

func GetDiaries(c echo.Context) error {
	var diaries = new(models.Diaries)
	models.Db.Find(diaries)
	return c.JSON(http.StatusAccepted, diaries)
}

func CreateDiary(c echo.Context) error {
	diary := new(models.Diary)
	if err := c.Bind(diary); err != nil {
		return err
	}
	models.Db.Create(diary)
	return c.Redirect(http.StatusMovedPermanently, "/diaries")
}

func GetDiary(c echo.Context) error {
	id := c.Param("id")
	diary := new(models.Diary)
	models.Db.First(diary, id)
	return c.JSON(http.StatusAccepted, diary)
}

func UpdateDiary(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}

func DeleteDiary(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}
