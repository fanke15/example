package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

const defaultJson = `

{"type":"page","title":"default","body":[{"type":"chart","config":{"xAxis":{"type":"category","data":["Mon","Tue","Wed","Thu","Fri","Sat","Sun"]},"yAxis":{"type":"value"},"series":[{"data":[820,932,901,934,1290,1330,1320],"type":"line"}]},"replaceChartOption":true,"id":"u:5897ec2e03fe","dataFilter":"","api":""}],"id":"u:c508a28b833b"}
`

func main() {
	r := gin.Default()

	r.Static("/assets", "./")

	tpl, err := ace.Load("default", "", &ace.Options{DynamicReload: true})
	if err != nil {
		fmt.Println("EOF:", err)
		return
	}
	r.GET("/", func(c *gin.Context) {
		_ = tpl.Execute(c.Writer, map[string]interface{}{
			"pageJson": defaultJson,
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
