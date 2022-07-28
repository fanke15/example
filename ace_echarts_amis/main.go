package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./")

	tpl, err := ace.Load("default", "", nil)
	if err != nil {
		fmt.Println("EOF:", err)
		return
	}

	r.GET("/", func(c *gin.Context) {
		_ = tpl.Execute(c.Writer, nil)
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
