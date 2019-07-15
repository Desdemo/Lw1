package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Query struct {
	Name   string   `json:"name"`
	Choose string   `json:"choose"`
	Pro    []string `json:"pro"`
}

func main() {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		var f []Query
		if err := c.BindJSON(&f); err != nil {
			return
		}
		c.IndentedJSON(http.StatusOK, f)
		fmt.Println(f)
	})
	r.Run(":4001")

}
