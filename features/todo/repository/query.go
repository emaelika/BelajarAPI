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

func (tq query) GetTodos(id uint) ([]todo.Todo, error) {
	var data []TodoModel

	if err := tq.connection.Where("user_id = ?", id).Find(&data).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var results []todo.Todo
	for _, val := range data {
		var result = todo.Todo{
			ID:        val.ID,
			UserID:    val.UserID,
			Kegiatan:  val.Kegiatan,
			Deskripsi: val.Deskripsi,
			Deadline:  val.Deadline,
		}
		results = append(results, result)
	}

	return results, nil
}
