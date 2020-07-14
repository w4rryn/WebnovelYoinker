package yoinker

//IYoinkerFactory defines an interface to crate a yoinker factory
type IYoinkerFactory interface {
	GetYoinker() IYoinkerManager
}

type yoinkerFactory struct {
	scraper IScrapingStrategy
	export  IExportStrategy
}

//GetYoinker creaes a new yoinker factory
func (y *yoinkerFactory) GetYoinker() IYoinkerManager {
	return &webnovelYoinker{
		scraper:  y.scraper,
		exporter: y.export,
	}
}

//NewYoinkerFactory creates a new factory instance
func NewYoinkerFactory(scraper IScrapingStrategy, exporter IExportStrategy) IYoinkerFactory {
	return &yoinkerFactory{
		scraper: scraper,
		export:  exporter,
	}
}
