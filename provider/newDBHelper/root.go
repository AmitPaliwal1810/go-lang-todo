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

//====================================== INSERT TODO ================================================================

func (dh DbHelper) InsertTodo(todo models.InsertTodo) error {
	insertQuery := `INSERT INTO todos (todo_value) VALUES($1)`

	_, err := dh.DB.Exec(insertQuery, todo.Todo)

	if err != nil {
		return err
	}

	return nil
}

//====================================== GET TODO =====================================================================

func (dh DbHelper) GetTodos() ([]models.Todos, error) {
	if dh.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	getQuery := `SELECT id, todo_value, is_completed, created_at FROM todos`
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

//======================================= UPDATE TODO =====================================================================

func (dh DbHelper) UpdateTodo(todo models.UpdateTodo) error {

	updateQuery := `UPDATE todos SET todo_value=$1, is_completed=$2 WHERE id=$3`

	_, err := dh.DB.Exec(updateQuery, todo.TodoValue, todo.IsCompleted, todo.Id)

	if err != nil {
		fmt.Println("got error while executing query")
		return err
	}

	return nil
}
