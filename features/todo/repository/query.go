package repository

import (
	"21-api/features/todo"
	"log"

	"gorm.io/gorm"
)

type query struct {
	connection *gorm.DB
}

func NewTodoQuery(db *gorm.DB) todo.TodoQuery {
	return &query{
		connection: db,
	}
}

func (tq query) AddTodo(newTodo todo.Todo) (todo.Todo, error) {
	var newData = TodoModel{
		Kegiatan:  newTodo.Kegiatan,
		Deskripsi: newTodo.Deskripsi,
		Deadline:  newTodo.Deadline,
		UserID:    newTodo.UserID,
	}
	err := tq.connection.Create(&newData).Error
	if err != nil {
		log.Println(err.Error())
		return todo.Todo{}, err
	}
	var result = todo.Todo{
		ID:        newData.ID,
		UserID:    newData.UserID,
		Kegiatan:  newData.Kegiatan,
		Deskripsi: newData.Deskripsi,
		Deadline:  newData.Deadline,
	}

	return result, nil
}
