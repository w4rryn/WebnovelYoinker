package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//IYoinkerManager Provides Functionality to yoink Webnovels and Webtoons
type IYoinkerManager interface {
	StartYoink(metadata book.Metadata, exportPath string) string
	GetAvailableVolumes(url string, website string) []book.Volume
}

//WebnovelYoinker scrapes webnovels and webtoons and exports them as epub or pdf
type webnovelYoinker struct {
	scraper  IScrapingStrategy
	exporter IExportStrategy
}

// StartYoink start yoinking the specified book
func (y *webnovelYoinker) StartYoink(metadata book.Metadata, path string) string {
	chapterChannel := make(chan book.Chapter, 5)
	go y.scraper.BeginScrape(metadata.ChapterURLs, chapterChannel)
	return y.exporter.Export(metadata, path, chapterChannel)
}

//GetAvailableVolumes get all available volumes of provided url
func (y *webnovelYoinker) GetAvailableVolumes(url string, website string) []book.Volume {
	return y.scraper.GetAvailableChapters(url)
}
