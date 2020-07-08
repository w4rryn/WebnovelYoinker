package yc

//ScrapingStrategy defines an interface to scrape from an arbitrary website
type ScrapingStrategy interface {
	// BeginScrape(metadata BookMetadata, chapterURLs []string) (*Volume, error)
	BeginScrape(chapterURLs []string, chapterChannel chan<- Chapter)
}
