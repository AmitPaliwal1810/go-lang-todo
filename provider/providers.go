package provider

import "github.com/AmitPaliwal1810/go-lang-todo/models"

type DBProvider interface {
	InsertTodo(todo models.InsertTodo) error
	GetTodos() ([]models.Todos, error)
	UpdateTodo(todo models.UpdateTodo) error
	// CreateUser(user models.CreateUser) error
	UserDbProvider
}

type UserDbProvider interface {
	CreateUser(user models.CreateUser) error
	GetAllUsers() ([]models.User, error)
	DeleteUser(user models.DeleteUser) error
	UpdateUser(user models.UpdateUser) error
	LoginUser(user models.Login) (models.UpdateUser, error)
}

type HelperFuncProvider interface {
	HashPassword(string) (string, error)
}
