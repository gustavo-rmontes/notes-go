package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	BodyContent string `json:"bodycontent"`
}

func getNotes(c *gin.Context) {
	dsn := "host=localhost user=gusdev password=14072003 dbname=notes_app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to db")
	}

	var someNotes Note

	db.Find(&someNotes)

	c.IndentedJSON(http.StatusOK, someNotes)
}

func createNote(c *gin.Context) {
	dsn := "host=localhost user=gusdev password=14072003 dbname=notes_app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to db")
	}

	var newNote Note

	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Note not binded")
		return
	}

	db.Create(&newNote)

	c.IndentedJSON(http.StatusCreated, "New note created")
}

func main() {
	dsn := "host=localhost user=gusdev password=14072003 dbname=notes_app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to db")
	}

	db.AutoMigrate(&Note{})

	router := gin.Default()

	router.GET("/notes", getNotes)

	router.POST("/notes", createNote)

	router.Run("localhost:8080")
}
