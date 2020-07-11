package yoinker

//MessageCallback prints messages to main thread
var MessageCallback func(s string)

func consoleLog(s string) {
	if MessageCallback != nil {
		MessageCallback(s)
	}
}

//WebnovelYoinker scrapes webnovels and webtoons and exports them as epub or pdf
type WebnovelYoinker struct {
	callback func(s string)
	Scraper  map[string]ScrapingStrategy
	Exporter map[string]ExportStrategy
}

// StartYoink start yoinking the specified book
func (y *WebnovelYoinker) StartYoink(metadata BookMetadata, path string) {
	chapterChannel := make(chan chapter, 5)
	go y.Scraper[metadata.WebsiteURL].BeginScrape(metadata.ChapterURLs, chapterChannel)
	y.Exporter[metadata.Format].Export(metadata, path, chapterChannel)
}
