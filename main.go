package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Status    string `json:"status"`
	Price     int    `json:"price"`
	Completed bool   `json:"description"`
}

var todos = []todo{
	{Id: "1", Author: "John", Status: "In Progress", Price: 100, Completed: false},
	{Id: "2", Author: "jado", Status: "Done", Price: 100, Completed: false},
	{Id: "3", Author: "poln", Status: "ok", Price: 100, Completed: false},
}

func DefoultGet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)

}

func Addcontext(c *gin.Context) {
	var newtodo todo
	if err := c.BindJSON(&newtodo); err != nil {
		return
	}
	todos = append(todos, newtodo)
	c.IndentedJSON(http.StatusCreated, newtodo)

}
func GetbyId(c *gin.Context) {
	id := c.Param("id")
	todo, err  := GettodobyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, *todo)
}

func GettodobyID(id string) (*todo, error) {

	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("todo not found baby")	
}

func toggleStatus(c *gin.Context){
	id := c.Param("id")
	todo, err  := GettodobyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	todo.Completed =!todo.Completed
	c.IndentedJSON(http.StatusOK, *todo)
}

func main() {

	router := gin.Default()

	router.GET("/todos", DefoultGet)
	router.POST("/todos", Addcontext)
	router.PATCH("/todos/:id", toggleStatus)
	router.GET("/todos/:id", GetbyId)
	router.Run("localhost:8080")

}
