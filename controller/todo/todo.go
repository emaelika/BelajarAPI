package todo

import (
	"21-api/helper"
	"21-api/middlewares"
	"21-api/model"
	"log"
	"net/http"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Model model.TodoModel
}

func (us *TodoController) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*golangjwt.Token)
		id, err := middlewares.ExtractId(token)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "harap login", nil))

		}
		var input TodoRequest
		err = c.Bind(&input)
		if err != nil {
			log.Println(err.Error())
			if strings.Contains(err.Error(), "unsupport") {

				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {

			log.Println(err.Error())
			// var message = []string{}
			// for _, val := range err.(validator.ValidationErrors) {
			// 	if val.Tag() == "required" {
			// 		message = append(message, fmt.Sprint(val.Field(), " wajib diisi"))
			// 	} else if val.Tag() == "min" {
			// 		message = append(message, fmt.Sprint(val.Field(), " minimal 10 digit"))
			// 	} else {
			// 		message = append(message, fmt.Sprint(val.Field(), " ", val.Tag()))
			// 	}
			// }
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim kurang sesuai", nil))
		}

		var processInput model.Todo
		processInput.Kegiatan = input.Kegiatan
		processInput.Deskripsi = input.Deskripsi
		processInput.Deadline = input.Deadline
		processInput.UserID = id

		result, err := us.Model.AddTodo(processInput) // ini adalah fungsi yang kita buat sendiri
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		var data TodoResponse

		data.Deadline = result.Deadline
		data.Deskripsi = result.Deskripsi
		data.Kegiatan = result.Kegiatan
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", data))
	}
}

// func (us *TodoController) UpdateTodo() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var hp = c.Param("hp")
// 		var input model.Todo
// 		err := c.Bind(&input)
// 		if err != nil {
// 			log.Println("masalah baca input:", err.Error())
// 			if strings.Contains(err.Error(), "unsupport") {
// 				return c.JSON(http.StatusUnsupportedMediaType,
// 					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
// 			}
// 			return c.JSON(http.StatusBadRequest,
// 				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
// 		}

// 		isFound := us.Model.CekTodo(hp)

// 		if !isFound {
// 			return c.JSON(http.StatusNotFound,
// 				helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
// 		}

// 		err = us.Model.Update(hp, input)

// 		if err != nil {
// 			log.Println("masalah database :", err.Error())
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
// 		}

// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "data berhasil di update", nil))
// 	}
// }

// func (us *TodoController) GetTodos() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		listTodo, err := us.Model.GetAllTodo()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}
// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", listTodo))
// 	}
// }

// func (us *TodoController) GetTodo() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var hp = c.Param("hp")
// 		result, err := us.Model.GetProfile(hp)

// 		if err != nil {
// 			if strings.Contains(err.Error(), "not found") {
// 				return c.JSON(http.StatusNotFound,
// 					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
// 			}
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}
// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
// 	}
// }
