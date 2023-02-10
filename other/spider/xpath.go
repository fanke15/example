package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

func main() {
	Demo()
}

func Demo() {
	doc, err := htmlquery.LoadURL("https://dune.com/fanke/postat/")
	nodes, err := htmlquery.QueryAll(doc, "/html/body/div/div/main/div[1]/div/section/div/div/article[1]/div/div/table")
	if err != nil {
		panic(`not a valid XPath expression.`)
	}

	fmt.Println(nodes[0].LastChild)
}
