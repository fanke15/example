package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	totalPageNum := int(math.Ceil((DateInterval("2006-01-02 15:04", "2020-11-03 00:00", "2022-07-12 00:00") + 1) / 25))

	var c, cancel = initDB()
	defer cancel()

	var respDodys = make([]*string, 0)

	// run task list
	sel := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > div > table > tbody"

	clickSel := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > ul > li:nth-child(6) > button"

	// clickSelOpt1 := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > ul > li:nth-child(5) > select > option:nth-child(7)"

	var clickOpt = make([]chromedp.Action, 0)
	clickOpt = append(clickOpt, chromedp.Navigate("https://dune.com/fanke/postat"))
	clickOpt = append(clickOpt, chromedp.WaitVisible(sel))

	for i := 0; i < totalPageNum; i++ {
		var res string
		if i > 0 {
			clickOpt = append(clickOpt, chromedp.Click(clickSel))
			clickOpt = append(clickOpt, chromedp.OuterHTML("html", &res))
		} else {
			clickOpt = append(clickOpt, chromedp.OuterHTML("html", &res))
		}
		respDodys = append(respDodys, &res)
	}

	if err := chromedp.Run(c, clickOpt...); err != nil {
		log.Fatalf("Failed getting body: %v", err)
	}

	log.Println("Body of duckduckgo.com starts with:")
	// log.Println(body)

	for index, v := range respDodys {
		dom, e := goquery.NewDocumentFromReader(strings.NewReader(*v))
		if e != nil {
			fmt.Println(1111, e)
		}
		dom.Find("tbody tr").Each(func(i int, sel1 *goquery.Selection) {
			sel1.Find("td").Each(func(j int, sel2 *goquery.Selection) {
				fmt.Println(index, i, j, sel2.Text())
			})
		})
	}
}

func initDB() (c context.Context, cancel context.CancelFunc) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	chromedp.Run(ctxt, make([]chromedp.Action, 0, 1)...)
	timeoutCtx, cancel := context.WithTimeout(ctxt, 30*time.Second)
	return timeoutCtx, cancel
}

// 计算两个日期间隔
func DateInterval(format, start, end string) float64 {
	startDate, err := time.ParseInLocation(format, start, time.Local)
	if err != nil {
		log.Println(err)
		return 0
	}
	endDate, err := time.ParseInLocation(format, end, time.Local)
	if err != nil {
		log.Println(err)
		return 0
	}
	return endDate.Sub(startDate).Hours() / 24
}
