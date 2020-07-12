package scrape

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//WuxiaWorldScraper scraping strategy for Wuxiaworld.com
type WuxiaWorldScraper struct {
}

// BeginScrape (metadata BookMetadata, chapterURLs []string) (*Volume, error)
func (w *WuxiaWorldScraper) BeginScrape(chapterURLs []string, chapterChannel chan<- book.Chapter) {
	panic("not implemented") // TODO: Implement
}

//GetAvailableChapters returns an array with all possible chapters
func (w *WuxiaWorldScraper) GetAvailableChapters(url string) []book.Volume {
	panic("not implemented") // TODO: Implement
}
