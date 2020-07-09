package yoinker

//NewYoinker Initializes a new Yoinker
func NewYoinker() YoinkManager {
	return &WebnovelYoinker{
		Scraper: map[string]ScrapingStrategy{
			"crimsonmagic": &crimsonmagicNovelScraper{},
		},
		Exporter: map[string]ExportStrategy{
			"epub": &epubExporter{},
		},
	}
}
