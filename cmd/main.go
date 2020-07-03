package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

var collector *colly.Collector

//chapter: https://www.crimsonmagic.me/joshikousei/jk-1-p/
func main() {
	collector = colly.NewCollector(
		colly.AllowedDomains("www.crimsonmagic.me"),
	)

	collector.OnHTML("div[class]", scraperCallback)

	collector.Visit("https://www.crimsonmagic.me/joshikousei/JK-1-1/")
}

//TODO add functionality to set chapters
func scraperCallback(e *colly.HTMLElement) {
	if e.Attr("class") != "entry-content" {
		return
	}
	fmt.Println(e.Text)
	content := e.ChildTexts("p")
	file, err := os.Create("dump.txt")
	check(err)

	defer file.Close()

	for lineNumber := range content {
		//TODO
		n, err := file.WriteString(content[lineNumber] + "\n")
		check(err)
		fmt.Printf("Wrote %d \n", n)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
