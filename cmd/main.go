package main

import (
	"fmt"

	export "github.com/lethal-bacon0/WebnovelYoinker/pkg/exportStrategies"
	scraping "github.com/lethal-bacon0/WebnovelYoinker/pkg/scrapingStrategies"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

func main() {
	chapters := []string{
		// "https://www.crimsonmagic.me/joshikousei/jk-1-p/",
		// "https://www.crimsonmagic.me/joshikousei/JK-1-1/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-2/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-3/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-4/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-5/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-6/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-7/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-8/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-9/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-10/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-11/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-12/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-13/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-14/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-15/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-16/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-17/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-e/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-a/",
		// "https://www.crimsonmagic.me/joshikousei/jk-1-ss/",
	}

	yoink := yoinker.WebnovelYoinker{
		Exporter: &export.EpubExporter{},
		Scraper: &scraping.CrimsonmagicNovelScraper{
			Callback: statusCallback,
		},
	}

	yoink.Scraper.SetMetadata("yuNS",
		"https://crimsonmagicme.files.wordpress.com/2018/08/cover1.jpg",
		"English",
		"I Shaved. Then I Brought a High School Girl Home. Volume 1",
		"2018")

	volume, err := yoink.Scraper.BeginScrape(chapters)
	if err != nil {
		panic(err)
	}
	yoink.Exporter.Export(volume)
}

func statusCallback(s string) {
	fmt.Println(s)
}

// c.volume = yoinker.Volume{
// 	Chapters: getChapters(chapterURLs, callback),
// 	Author:   "yuNS",
// 	Cover:    "https://crimsonmagicme.files.wordpress.com/2018/08/cover1.jpg",
// 	Language: "English",
// 	Title:    "I Shaved. Then I Brought a High School Girl Home. Volume 1",
// 	Year:     "2018",
// }
