package handler

import (
	"21-api/features/file"
	"fmt"
	"io"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	s file.Service
}

func New(s file.Service) file.Handler {
	return &Handler{
		s: s,
	}
}

func (at *Handler) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.FormValue("id")
		// name := c.FormValue("name")
		// email := c.FormValue("email")

		//-----------
		// Read file
		//-----------

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dir := "uploads"
		dirUser := dir + "/" + id
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println(dir, "does not exist")
			os.Mkdir(dir, 0644)

		} else {
			fmt.Println("The provided directory named", dir, "exists")
		}
		if _, err := os.Stat(dirUser); os.IsNotExist(err) {
			fmt.Println(dirUser, "does not exist")
			os.Mkdir(dirUser, 0644)

		} else {
			fmt.Println("The provided directory named", dir, "exists")
		}
		// Destination
		dst, err := os.Create(dirUser + "/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "success",
			"data": dirUser + "/" + file.Filename})
	}
}

func (ct *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		filename := c.Param("filename")
		id := c.Param("id")
		if filename == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "filepath gak ada"})
		}

		return c.File("uploads/" + id + "/" + filename)
	}
}
