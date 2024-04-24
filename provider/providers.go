package provider

import "github.com/AmitPaliwal1810/go-lang-todo/models"

type DBProvider interface {
	InsertTodo(todo models.InsertTodo) error
	GetTodos() ([]models.Todos, error)
	UpdateTodo(todo models.UpdateTodo) error
	CreateUser(user models.CreateUser) error
}
