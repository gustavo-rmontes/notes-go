package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Subtitle    string `json: "subtitle"`
	BodyContent string `json: "bodycontent"`
}

var someNotes = []Note{{ID: 1, Title: "Viagem a Porto Rico", Subtitle: "Explorando cavernas", BodyContent: "Por cinco dias, adentrei as cavernas..."}, {ID: 2, Title: "Trabalho malfeito", Subtitle: "Não volto a trabalhar com eles", BodyContent: "Descobri que ao..."}}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, someNotes)
}

func createNote(c *gin.Context) {
	var newNote Note

	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Note not binded")
		return
	}

	someNotes = append(someNotes, newNote)
	c.IndentedJSON(http.StatusCreated, "New note created")
}

func main() {
	router := gin.Default()

	router.GET("/notes", getNotes)

	router.POST("/notes", createNote)

	router.Run("localhost:8080")
}
