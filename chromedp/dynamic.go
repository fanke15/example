package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	chromedp.Run(ctxt, make([]chromedp.Action, 0, 1)...)

	timeoutCtx, cancel := context.WithTimeout(ctxt, 30*time.Second)
	defer cancel()

	// run task list
	sel := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > div > table > tbody"

	clickSel := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > ul > li:nth-child(6) > button"

	// clickSelOpt1 := "#__next > div > main > div:nth-child(1) > div > section > div > div > article:nth-child(1) > div > ul > li:nth-child(5) > select > option:nth-child(7)"
	var body string
	var clickOpt = make([]chromedp.Action, 0)
	clickOpt = append(clickOpt, chromedp.Navigate("https://dune.com/fanke/postat"))
	clickOpt = append(clickOpt, chromedp.WaitVisible(sel))
	for i := 0; i < 30; i++ {
		clickOpt = append(clickOpt, chromedp.Click(clickSel))
	}
	clickOpt = append(clickOpt, chromedp.OuterHTML("html", &body))

	for i := 0; i < 30; i++ {
		clickOpt = append(clickOpt, chromedp.Click(clickSel))
	}

	if err := chromedp.Run(timeoutCtx, clickOpt...); err != nil {
		log.Fatalf("Failed getting body: %v", err)
	}

	log.Println("Body of duckduckgo.com starts with:")
	// log.Println(body)

	dom, e := goquery.NewDocumentFromReader(strings.NewReader(body))
	if e != nil {
		fmt.Println(1111, e)
	}
	dom.Find("tbody tr").Each(func(i int, sel1 *goquery.Selection) {
		sel1.Find("td").Each(func(j int, sel2 *goquery.Selection) {
			fmt.Println(i, j, sel2.Text())
		})
	})
}

func GetDynamicData() {
	// 使用Chromedp获取动态数据的html

	// 使用goquery解析html代码 获取数据
}
