package scrape

import (
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//wuxiaWorldScraper scraping strategy for Wuxiaworld.com
type wuxiaWorldScraper struct {
}

// BeginScrape(metadata BookMetadata, chapterURLs []string) (*Volume, error)
func (w *wuxiaWorldScraper) ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter {
	panic("Not implemented") // TODO: implement  this
}

//GetAwvailableChapters returns an array with all possible chapters
func (w *wuxiaWorldScraper) GetAvailableChapters(url string) []book.Volume {
	panic("Not implemented") // TODO: implement  this
}

//NewWuxiaScraper creates a new wuxia scraper strategy
func NewWuxiaScraper() yoinker.IScrapingStrategy {
	return &wuxiaWorldScraper{}
}
