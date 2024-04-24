package main

import (
	"fmt"
	"log"      // log message to the console
	"net/http" // use to build http server and client

	"github.com/AmitPaliwal1810/go-lang-todo/handler"
	"github.com/go-chi/chi/v5" // this is for chi router
	// this is for middleware
)

const portNum string = ":8080"

func InfoPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info page")
}

func main() {
	r := chi.NewRouter()
	srv := handler.ServerInit()

	// r.Use(middleware.Logger)
	// r.Get("/", handler.GetTodo)

	r.Post("/dummy", handler.DummyApi)
	r.Post("/insertTodo", srv.InsertTodo)
	r.Get("/getTodo", srv.GetTodo)
	r.Put("/update-todo", srv.UpdateTodo)

	// user routes

	r.Post("/create-user", srv.CreateUser)

	log.Println("Starting HTTP server")
	http.HandleFunc("/info", InfoPage)
	log.Println("server is starting on", portNum)

	err := http.ListenAndServe(portNum, r)

	if err != nil {
		log.Fatal(err)
	}
}
