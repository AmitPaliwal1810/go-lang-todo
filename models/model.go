package models

type Dummy struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type InsertTodo struct {
	Todo        string `json:"todo"`
	IsCompleted bool   `json:"isCompleted"`
}

type Todos struct {
	ID          string `json:"id" db:"id"`
	TodoValue   string `json:"todoValue" db:"todovalue"`
	IsCompleted bool   `json:"isCompleted" db:"iscompleted"`
	CreateAt    string `json:"createdAt" db:"createat"`
}

type UpdateTodo struct {
	Id          string `json:"id" db:"id"`
	TodoValue   string `json:"todoValue" db:"todovalue"`
	IsCompleted bool   `json:"isCompleted" db:"iscompleted"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
