package yoinker

import (
	yc "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/yoinkerCore"
	yoinkerexport "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/yoinkerExport"
	yoinkerscrape "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/yoinkerScrape"
)

//NewYoinker Initializes a new Yoinker
func NewYoinker() YoinkManager {
	return &yc.WebnovelYoinker{
		Scraper: map[string]yc.ScrapingStrategy{
			"crimsonmagic": &yoinkerscrape.CrimsonmagicNovelScraper{},
		},
		Exporter: map[string]yc.ExportStrategy{
			"epub": &yoinkerexport.EpubExporter{},
		},
	}
}
