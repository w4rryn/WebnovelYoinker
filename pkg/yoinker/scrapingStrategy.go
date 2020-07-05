package yoinker

//ScrapingStrategy defines an interface to scrape from an arbitrary website
type ScrapingStrategy interface {
	BeginScrape(chapterURLs []string) (*Volume, error)
	SetMetadata(author, coverURL, language, title, year string)
}
