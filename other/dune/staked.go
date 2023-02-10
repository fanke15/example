package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
	"time"
)

func main() {
	GetStakedHistory()
}

// 通过爬虫获取总质押数据
func GetStakedHistory() {
	spiderC := spiderReq()

	spiderC.Limit(&colly.LimitRule{
		Delay: 30 * time.Second,
	})

	sel := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > div > table > tbody"
	spiderC.OnHTML(sel, func(e *colly.HTMLElement) {
		time.Sleep(10 * time.Second)
		fmt.Println(1111, e.Text)
	})

	if err := spiderC.Visit("https://dune.com/fanke/postat"); err != nil {
		fmt.Println(err)
	}

	spiderC.Wait()

	return

}

//---------------------------内部私有方法---------------------------//

const beaconApiDomain = "dune.com"

// spider请求封装
func spiderReq() *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(beaconApiDomain), // 限定爬取数据源
		colly.UserAgent("xy"),
	)
}

// 提取字符串中的浮点数
func getFiguresByStr(str string) string {
	re := regexp.MustCompile("[0-9]\\d*\\.?\\d*")
	return strings.Join(re.FindAllString(str, -1), "")
}
