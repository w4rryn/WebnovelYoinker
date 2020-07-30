package factories

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/scrape"
)

//GetScraper gets scraper instance of matching value
func GetScraper(scraper string) (yoinker.IScrapingStrategy, error) {
	switch scraper {
	case string(scrape.CRIMSON):
		return scrape.NewCrimsonmagicScraper(), nil
	case string(scrape.WUXIA):
		return scrape.NewWuxiaScraper(), nil
	}

	return nil, errors.New("Source not supported")
}
