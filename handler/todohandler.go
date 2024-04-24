package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
	"github.com/AmitPaliwal1810/go-lang-todo/provider/database"
	"github.com/AmitPaliwal1810/go-lang-todo/provider/newDBHelper"
)

// server structure for serverInit

type Response struct {
	Message string `json:"message"`
}

func ServerInit() *Server {
	DB, err := database.DBConnection(
		"localhost",
		"5433",
		"todo",
		"postgres",
		"1234",
		database.SSLModeDisable)

	if err != nil {
		fmt.Printf("Failed to initialize and migrate database with error: %+v", err)
	}
	fmt.Print("Database connected")

	// initializing new dbHelper, repository level;

	dbHelper := newDBHelper.NewDbHelperProvider(DB)

	return &Server{
		DbHelper: dbHelper,
	}
}

func DummyApi(w http.ResponseWriter, r *http.Request) {
	var user models.Dummy
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println("error aa gai hai.")
		return
	}

	fmt.Print("user", user)

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		fmt.Println("error aa gai hai. fir se")
		return
	}

}

func (srv *Server) InsertTodo(w http.ResponseWriter, r *http.Request) {
	var insertTodo models.InsertTodo

	err := json.NewDecoder(r.Body).Decode(&insertTodo)

	if err != nil {
		fmt.Print("getting error while decoding insertTodo function", err)
		return
	}

	err = srv.DbHelper.InsertTodo(insertTodo)

	if err != nil {
		fmt.Print("error insserting todo server database", err)
		return
	}

	_ = json.NewEncoder(w).Encode("success")

}

func (srv *Server) GetTodo(w http.ResponseWriter, r *http.Request) {

	data, err := srv.DbHelper.GetTodos()

	if err != nil {
		fmt.Print("error getting todo server database", err)
		return
	}

	fmt.Println("showing data", err)

	_ = json.NewEncoder(w).Encode(data)

}

func (srv *Server) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo models.UpdateTodo

	err := json.NewDecoder(r.Body).Decode(&updateTodo)

	if err != nil {
		fmt.Println("getting error while decoding updateTodo func", err)
	}

	err = srv.DbHelper.UpdateTodo(updateTodo)

	if err != nil {
		fmt.Println("error while call updateTodo from root", err)
		return
	}

	response := Response{Message: "success"}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
