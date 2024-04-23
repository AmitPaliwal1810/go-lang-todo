package newDBHelper

import (
	"fmt"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
	"github.com/AmitPaliwal1810/go-lang-todo/provider"
	"github.com/jmoiron/sqlx"
)

type DbHelper struct {
	DB *sqlx.DB
}

func NewDbHelperProvider(db *sqlx.DB) provider.DBProvider {
	return DbHelper{
		DB: db,
	}
}

func (dh DbHelper) InsertTodo(todo models.InsertTodo) error {
	insertQuery := `INSERT INTO todos (todovalue, iscompleted) VALUES($1,$2)`

	// _, err := dh.DB.Exec(insertQuery, args)
	_, err := dh.DB.Exec(insertQuery, todo.Todo, todo.IsCompleted)

	if err != nil {
		return err
	}

	return nil
}

func (dh DbHelper) GetTodos() ([]models.Todos, error) {
	if dh.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	getQuery := `SELECT id, todovalue, iscompleted, createat FROM todos`
	var todos []models.Todos
	err := dh.DB.Select(&todos, getQuery)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	fmt.Println("todo-data", todos)

	if len(todos) == 0 {
		fmt.Println("No todos found.")
	}

	return todos, nil
}

func (dh DbHelper) UpdateTodo(todo models.UpdateTodo) error {

	updateQuery := `UPDATE todos SET todovalue=$1, iscompleted=$2 WHERE id=$3`

	_, err := dh.DB.Exec(updateQuery, todo.TodoValue, todo.IsCompleted, todo.Id)

	if err != nil {
		fmt.Println("got error while executing query")
		return err
	}

	return nil
}
