package main

import (
	"github.com/fanke15/example/g2a/middleware"
	"github.com/fanke15/example/g2a/pages"
	"github.com/gin-gonic/gin"

	"github.com/buger/jsonparser"
)

func main() {
	router()
}

func router() {
	engine := gin.New()
	engine.Static("assets", "./static")
	engine.StaticFile("favicon.ico", "./static/img/favicon.ico")

	web := engine.Group("web")
	{
		web.GET("init1", middleware.InitAce(map[string]interface{}{
			"title": "init 1",
			"conf":  pages.Page1,
		}))

		v, _ := jsonparser.Set([]byte(pages.Page2), []byte(`
{
    "name": "zkf",
    "age": 666666
}
`), "data")
		web.GET("init2", middleware.InitAce(map[string]interface{}{
			"title": "init 2",
			"conf":  string(v),
		}))

		web.GET("init3", middleware.InitAce(map[string]interface{}{
			"title": "init 3",
			"conf":  string(pages.Page3),
		}))
	}

	engine.Run(":8080")
}
