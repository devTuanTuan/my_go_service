package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	r := gin.Default()

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "task 1",
		Description: "task 1 description",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsData, err := json.Marshal(item)

	if err != nil {
		log.Fatalln(err)
	}

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("")
			items.GET("")
			items.GET("/:id")
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": string(jsData),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
