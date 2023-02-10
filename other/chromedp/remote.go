package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background())
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	var body string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate("https://duckduckgo.com"),
		chromedp.WaitVisible("#logo_homepage_link"),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		log.Fatalf("Failed getting body of duckduckgo.com: %v", err)
	}

	log.Println("Body of duckduckgo.com starts with:")
	log.Println(body[0:100])
}
