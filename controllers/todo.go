package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/models"
	"html/template"
	"net/http"
)

var (
	id       int
	task     string
	assignee string
	deadline string
	status   int
	view     = template.Must(template.ParseFiles("./views/index.html"))
	database = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM go_task`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &task, &assignee, &deadline, &status)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:       id,
			Task:     task,
			Assignee: assignee,
			Deadline: deadline,
			Status:   status,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	task := r.FormValue("task")
	assignee := r.FormValue("assignee")
	deadline := r.FormValue("deadline")

	_, err := database.Exec(`INSERT INTO go_task (task,assignee,deadline) VALUE (?,?,?)`, task, assignee, deadline)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	_, err := database.Exec(`DELETE FROM go_task WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Status(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE go_task SET status = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
