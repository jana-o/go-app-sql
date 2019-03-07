package handlers

import (
	"code/go-app-sql/models"
	"database/sql"
	"io"
	"os"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetPhotos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPhotos(db))
	}
}

func UpdatePhoto(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var photo models.Photo

		c.Bind(&photo)

		index, _ := strconv.Atoi(c.Param("index"))

		id, err := models.UpdatePhoto(db, index, photo.Src)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"affected": id,
			})
		}

		return err
	}
}

func UploadPhoto(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var photo models.Photo

		c.Bind(&photo)

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		defer src.Close()

		fileName := file.Filename
		filePath := "./public/uploads/" + fileName
		fsrc := "http://127.0.0.1:8080/uploads/" + fileName

		dst, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, models.UploadPhoto(db, fsrc))
	}
}
