package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
	"log"
)

func InitAce(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		tpl, err := ace.Load("./middleware/index", "", &ace.Options{DynamicReload: true})
		if err != nil {
			log.Println("err:", err)
			return
		}
		tpl.Execute(c.Writer, data)
	}
}
