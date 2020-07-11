package factory

import (
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/export"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/scrape"
)

//NewYoinker Initializes a new Yoinker
func NewYoinker() yoinker.YoinkManager {
	return &yoinker.WebnovelYoinker{
		Scraper: map[string]yoinker.ScrapingStrategy{
			"crimsonmagic": &scrape.CrimsonmagicNovelScraper{},
		},
		Exporter: map[string]yoinker.ExportStrategy{
			"epub": &export.EpubExporter{},
		},
	}
}
