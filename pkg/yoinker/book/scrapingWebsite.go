package book

//ScrapingWebsite declares supported scraping websites
type ScrapingWebsite string

const (
	//CRIMSON scrapes from crimsonmagic
	CRIMSON ScrapingWebsite = "crimsonmagic"
	//WUXIA scrapes from wuxia
	WUXIA ScrapingWebsite = "wuxia"
	//WEBTOON Scxrapes from Webtoon
	WEBTOON ScrapingWebsite = "webtoon"
)
