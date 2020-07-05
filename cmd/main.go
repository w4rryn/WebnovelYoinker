package main

import (
	"fmt"

	export "github.com/lethal-bacon0/WebnovelYoinker/pkg/exportStrategies"
	scraping "github.com/lethal-bacon0/WebnovelYoinker/pkg/scrapingStrategies"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

func main() {
	// volume1 := []string{
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-p/",
	// 	"https://www.crimsonmagic.me/joshikousei/JK-1-1/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-2/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-3/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-4/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-5/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-6/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-7/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-8/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-9/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-10/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-11/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-12/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-13/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-14/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-15/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-16/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-17/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-e/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-a/",
	// 	"https://www.crimsonmagic.me/joshikousei/jk-1-ss/",
	// }

	konosuba := []string{
		"https://www.crimsonmagic.me/archive/gifting-1-p/",
		"https://www.crimsonmagic.me/archive/gifting-1-1/",
		"https://www.crimsonmagic.me/archive/gifting-1-2/",
		"https://www.crimsonmagic.me/archive/gifting-1-3/",
		"https://www.crimsonmagic.me/archive/gifting-1-4/",
		"https://www.crimsonmagic.me/archive/gifting-1-e/",
	}
	yoink := yoinker.WebnovelYoinker{
		Exporter: &export.EpubExporter{
			Callback: statusCallback,
		},
		Scraper: &scraping.CrimsonmagicNovelScraper{
			Callback: statusCallback,
		},
	}

	scrapeVolume1, err := yoink.Scraper.BeginScrape(yoinker.BookMetadata{
		Author:   "Natsume Akatsuki",
		Cover:    "https://i.imgur.com/eLIkho2.png",
		Language: "English",
		Title:    "Konosuba Volume 1",
		Year:     "2013",
	}, konosuba)
	if err != nil {
		panic(err)
	}
	yoink.Exporter.Export(scrapeVolume1)
}

func statusCallback(s string) {
	fmt.Println(s)
}
