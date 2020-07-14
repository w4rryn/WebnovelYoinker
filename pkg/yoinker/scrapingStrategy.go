package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//IScrapingStrategy defines an interface to scrape from an arbitrary website
type IScrapingStrategy interface {
	// BeginScrape(metadata BookMetadata, chapterURLs []string) (*Volume, error)
	ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter

	//GetAvailableChapters returns an array with all possible chapters
	GetAvailableChapters(url string) []book.Volume
}
