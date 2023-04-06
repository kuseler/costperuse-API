package main

import (
	"CostPerUse/DBInteraction"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	type item struct {
		ID          int       `json:"id"`
		name        string    `json:"name"`
		price       float64   `json:"price"`
		date_bought time.Time `json:"date_bought"`
		notes       string    `json:"notes"`
		category    int       `json:"category"`
	}

var items []item
*/

func main() {
	// get it by env var
	connStr := "postgresql://postgres:docker@localhost/cost_per_use_tracker?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// rows := getResAsJson(db, "SELECT * from items;")
	fmt.Println(DBInteraction.GetQueryResultAsItemList(db, "SELECT * from items;"))

	// generic boilerplate, only a placeholder for now
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
