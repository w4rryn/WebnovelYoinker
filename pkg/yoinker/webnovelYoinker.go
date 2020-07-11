package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//WebnovelYoinker scrapes webnovels and webtoons and exports them as epub or pdf
type WebnovelYoinker struct {
	callback func(s string)
	Scraper  map[string]ScrapingStrategy
	Exporter map[string]ExportStrategy
}

// StartYoink start yoinking the specified book
func (y *WebnovelYoinker) StartYoink(metadata book.Metadata, path string) {
	chapterChannel := make(chan book.Chapter, 5)
	go y.Scraper[metadata.WebsiteURL].BeginScrape(metadata.ChapterURLs, chapterChannel)
	y.Exporter[metadata.Format].Export(metadata, path, chapterChannel)
}
