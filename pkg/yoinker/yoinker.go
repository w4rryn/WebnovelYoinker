package yoinker

//WebnovelYoinker scrapes webnovels and webtoons and exports them as epub or pdf
type WebnovelYoinker struct {
	callback func(s string)
	Scraper  ScrapingStrategy
	Exporter ExportStrategy
}
