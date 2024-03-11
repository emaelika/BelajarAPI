package todo

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController interface {
	AddTodo() echo.HandlerFunc
}

type TodoService interface {
	AddTodo(pemilik *jwt.Token, kegiatanBaru Todo) (Todo, error)
}

type TodoQuery interface {
	AddTodo(newData Todo) (Todo, error)
}

type Todo struct {
	ID        uint
	UserID    uint
	Kegiatan  string
	Deskripsi string
	Deadline  string
}
