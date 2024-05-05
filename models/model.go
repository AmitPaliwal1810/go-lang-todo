package models

type Dummy struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type InsertTodo struct {
	Todo string `json:"todo"`
}

type Todos struct {
	ID          string `json:"id" db:"id"`
	TodoValue   string `json:"todoValue" db:"todo_value"`
	IsCompleted bool   `json:"isCompleted" db:"is_completed"`
	CreateAt    string `json:"createdAt" db:"created_at"`
}

type UpdateTodo struct {
	Id          string `json:"id" db:"id"`
	TodoValue   string `json:"todoValue" db:"todo_value"`
	IsCompleted bool   `json:"isCompleted" db:"is_completed"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	CreateAt string `json:"created_at" db:"created_at"`
}

type DeleteUser struct {
	Id string `json:"id"`
}

type UpdateUser struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Login struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}