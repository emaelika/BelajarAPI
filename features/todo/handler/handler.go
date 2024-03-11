package handler

import (
	"21-api/features/todo"
	"21-api/helper"
	"log"
	"net/http"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type controller struct {
	s todo.TodoService
}

func NewHandler(service todo.TodoService) todo.TodoController {
	return &controller{
		s: service,
	}
}
func (us *controller) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := c.Get("user").(*golangjwt.Token)

		var input TodoRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println(err.Error())
			if strings.Contains(err.Error(), "unsupport") {

				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		var processInput todo.Todo
		processInput.Kegiatan = input.Kegiatan
		processInput.Deskripsi = input.Deskripsi
		processInput.Deadline = input.Deadline

		result, err := us.s.AddTodo(token, processInput) // ini adalah fungsi yang kita buat sendiri
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

// func (us *TodoController) GetTodos() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := c.Get("user").(*golangjwt.Token)
// 		id, err := middlewares.ExtractId(token)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "harap login", nil))

// 		}
// 		listTodo, err := us.Model.GetTodos(id)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}
// 		var result []TodoResponse
// 		for _, val := range listTodo {
// 			var data TodoResponse
// 			data.Deadline = val.Deadline
// 			data.Deskripsi = val.Deskripsi
// 			data.Kegiatan = val.Kegiatan

// 			result = append(result, data)
// 		}

// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
// 	}
// }

// func (us *TodoController) GetTodo() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := c.Get("user").(*golangjwt.Token)
// 		id, err := middlewares.ExtractId(token)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "harap login", nil))

// 		}
// 		idTodo, _ := strconv.Atoi(c.Param("id"))

// 		val, err := us.Model.GetTodo(uint(idTodo))
// 		if err != nil {
// 			if strings.Contains(err.Error(), "not found") {
// 				return c.JSON(http.StatusNotFound,
// 					helper.ResponseFormat(http.StatusNotFound, "tidak ditemukan to do", nil))
// 			}
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}
// 		var result TodoResponse
// 		if val.UserID != id {
// 			return c.JSON(http.StatusUnauthorized,
// 				helper.ResponseFormat(http.StatusUnauthorized, "anda tidak bisa mengakses item ini", nil))
// 		}
// 		result.Deadline = val.Deadline
// 		result.Deskripsi = val.Deskripsi
// 		result.Kegiatan = val.Kegiatan

// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
// 	}
// }

// func (us *TodoController) UpdateTodo() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Ambil ID
// 		token := c.Get("user").(*golangjwt.Token)
// 		id, err := middlewares.ExtractId(token)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "harap login", nil))

// 		}
// 		idTodo, _ := strconv.Atoi(c.Param("id"))

// 		// Bind Input
// 		var input TodoUpdateRequest
// 		err = c.Bind(&input)
// 		if err != nil {
// 			log.Println(err.Error())
// 			if strings.Contains(err.Error(), "unsupport") {

// 				return c.JSON(http.StatusUnsupportedMediaType,
// 					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
// 			}
// 			return c.JSON(http.StatusBadRequest,
// 				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
// 		}

// 		// Validasi
// 		validate := validator.New(validator.WithRequiredStructEnabled())
// 		err = validate.Struct(input)

// 		if err != nil {

// 			log.Println(err.Error())
// 			var message = []string{}
// 			for _, val := range err.(validator.ValidationErrors) {

// 				message = append(message, fmt.Sprint("error pada ", val.Field()))
// 			}
// 			return c.JSON(http.StatusBadRequest,
// 				helper.ResponseFormat(http.StatusBadRequest, message, nil))
// 		}

// 		// cek kesamaan ID User
// 		val, err := us.Model.GetTodo(uint(idTodo))
// 		if err != nil {
// 			if strings.Contains(err.Error(), "not found") {
// 				return c.JSON(http.StatusNotFound,
// 					helper.ResponseFormat(http.StatusNotFound, "tidak ditemukan to do", nil))
// 			}
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}

// 		if val.UserID != id {
// 			return c.JSON(http.StatusUnauthorized,
// 				helper.ResponseFormat(http.StatusUnauthorized, "anda tidak bisa mengakses item ini", nil))
// 		}

// 		// parsing
// 		var processInput model.Todo
// 		if input.Kegiatan != "" {
// 			processInput.Kegiatan = input.Kegiatan
// 		} else {
// 			processInput.Kegiatan = val.Kegiatan
// 		}
// 		if input.Deskripsi != "" {
// 			processInput.Deskripsi = input.Deskripsi
// 		} else {
// 			processInput.Deskripsi = val.Deskripsi
// 		}
// 		if input.Deadline != "" {
// 			processInput.Deadline = input.Deadline
// 		} else {
// 			processInput.Deadline = val.Deadline
// 		}
// 		processInput.ID = val.ID

// 		// Input
// 		update, err := us.Model.UpdateTodo(processInput) // ini adalah fungsi yang kita buat sendiri
// 		if err != nil {
// 			log.Println(err.Error())
// 			return c.JSON(http.StatusInternalServerError,
// 				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
// 		}

// 		var result TodoResponse
// 		result.Deadline = update.Deadline
// 		result.Deskripsi = update.Deskripsi
// 		result.Kegiatan = update.Kegiatan

// 		return c.JSON(http.StatusOK,
// 			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
// 	}
// }
