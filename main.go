package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	ITEM      string `json:"item"`
	COMPLETED bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", ITEM: "book", COMPLETED: true},
	{ID: "2", ITEM: "read", COMPLETED: true},
	{ID: "3", ITEM: "eat", COMPLETED: false},
}

func getTodos(context *gin.Context) { //first context include incoming http request
	context.IndentedJSON(http.StatusOK, todos) //todos struct convert to json
}
func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil { /* just json is acceptable*/
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.COMPLETED = !todo.COMPLETED
	context.IndentedJSON(http.StatusOK, todo)
}
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func main() {
	router := gin.Default()
	router.GET("/todos/:id", getTodos) //return the client
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todis", addTodo)
	router.Run("localhost:8080")
}
