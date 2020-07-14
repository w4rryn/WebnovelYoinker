package scrape

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//GetScraper gets scraper instance of matching value
func GetScraper(scraper book.ScrapingWebsite) (yoinker.IScrapingStrategy, error) {
	switch scraper {
	case book.CRIMSON:
		return NewCrimsonmagicScraper(), nil
	}

	return nil, errors.New("Source not supported")
}
