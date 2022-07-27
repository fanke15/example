package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

func main() {
	r := gin.Default()

	tpl, err := ace.Load("default", "", nil)
	if err != nil {
		fmt.Println("EOF:", err)
		return
	}

	r.GET("/", func(c *gin.Context) {
		tpl.Execute(c.Writer, nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
