package factories

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/scrape"
)

//GetScraper gets scraper instance of matching value
func GetScraper(scraper book.ScrapingWebsite) (yoinker.IScrapingStrategy, error) {
	switch scraper {
	case book.CRIMSON:
		return scrape.NewCrimsonmagicScraper(), nil
	case book.WUXIA:
		return scrape.NewWuxiaScraper(), nil
	case book.FULLBOOKEDTLS:
		return scrape.NewFullyBookedTLSSCraper(), nil
	}

	return nil, errors.New("Source not supported")
}
