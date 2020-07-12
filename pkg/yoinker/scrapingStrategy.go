package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//ScrapingStrategy defines an interface to scrape from an arbitrary website
type ScrapingStrategy interface {
	// BeginScrape(metadata BookMetadata, chapterURLs []string) (*Volume, error)
	BeginScrape(chapterURLs []string, chapterChannel chan<- book.Chapter)

	//GetAvailableChapters returns an array with all possible chapters
	GetAvailableChapters(url string) []book.Volume
}
