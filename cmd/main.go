package main

import (
	"fmt"
	"sync"

	export "github.com/lethal-bacon0/WebnovelYoinker/pkg/exportStrategies"
	scraping "github.com/lethal-bacon0/WebnovelYoinker/pkg/scrapingStrategies"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

func main() {
	var waiter sync.WaitGroup
	waiter.Add(2)

	go func() {
		volume1 := []string{
			"https://www.crimsonmagic.me/joshikousei/jk-1-p/",
			"https://www.crimsonmagic.me/joshikousei/JK-1-1/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-2/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-3/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-4/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-5/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-6/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-7/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-8/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-9/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-10/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-11/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-12/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-13/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-14/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-15/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-16/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-17/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-e/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-a/",
			"https://www.crimsonmagic.me/joshikousei/jk-1-ss/",
		}
		yoink := newYoinker()
		meta := yoinker.BookMetadata{
			Author:   "Shimesaba",
			Cover:    "https://crimsonmagicme.files.wordpress.com/2018/08/cover1.jpg",
			Language: "English",
			Title:    "I Shaved. Then I Brought a High School Girl Home Volume 1",
			Year:     "2018",
		}
		chapterChannel := make(chan yoinker.Chapter, 5)
		go yoink.Scraper.BeginScrape(volume1, chapterChannel)
		yoink.Exporter.Export(meta, "", chapterChannel)
		waiter.Done()
	}()

	go func() {
		konosuba := []string{
			"https://www.crimsonmagic.me/archive/gifting-1-p/",
			"https://www.crimsonmagic.me/archive/gifting-1-1/",
			"https://www.crimsonmagic.me/archive/gifting-1-2/",
			"https://www.crimsonmagic.me/archive/gifting-1-3/",
			"https://www.crimsonmagic.me/archive/gifting-1-4/",
			"https://www.crimsonmagic.me/archive/gifting-1-e/",
		}
		yoink := newYoinker()
		meta := yoinker.BookMetadata{
			Author:   "Natsume Akatsuki",
			Cover:    "https://i.imgur.com/eLIkho2.png",
			Language: "English",
			Title:    "Konosuba Volume 1",
			Year:     "2013",
		}
		chapterChannel := make(chan yoinker.Chapter, 5)
		go yoink.Scraper.BeginScrape(konosuba, chapterChannel)
		yoink.Exporter.Export(meta, "", chapterChannel)
		waiter.Done()
	}()

	waiter.Wait()
}

func newYoinker() *yoinker.WebnovelYoinker {
	return &yoinker.WebnovelYoinker{
		Exporter: &export.EpubExporter{
			PrintCallback: printCallback,
		},
		Scraper: &scraping.CrimsonmagicNovelScraper{
			PrintCallback: printCallback,
		},
	}
}

func printCallback(s string) {
	fmt.Println(s)
}
