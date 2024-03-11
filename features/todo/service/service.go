package service

import (
	"21-api/features/todo"
	"21-api/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	tq todo.TodoQuery
	v  *validator.Validate
}

func NewTodoService(query todo.TodoQuery) todo.TodoService {
	return &service{
		tq: query,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (ts service) AddTodo(pemilik *jwt.Token, kegiatanBaru todo.Todo) (todo.Todo, error) {
	userID, err := middlewares.ExtractId(pemilik)
	if err != nil {
		log.Println(err.Error())
		return todo.Todo{}, err
	}

	err = ts.v.Struct(&kegiatanBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return todo.Todo{}, errors.New("data tidak valid")
	}

	kegiatanBaru.UserID = userID
	result, err := ts.tq.AddTodo(kegiatanBaru)
	if err != nil {
		log.Println("query error", err.Error())
		return todo.Todo{}, err
	}
	return result, nil

}
