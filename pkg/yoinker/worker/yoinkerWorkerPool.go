package worker

import (
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/factory"
)

//ScrapeWorkerpool privides functionality to create a background worker to export multiple books
type ScrapeWorkerpool interface {
	StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan book.Metadata, resultChannel chan string, exportPath string)
	BeginMultiConvert(books []book.Metadata, numberOfWorkers int, outputPath string)
}

//YoinkWorkerPool workerpool to scrape and export books
type yoinkWorkerPool struct {
}

//StartScrapeWorkerPool start the yoinker worker pool
func (y *yoinkWorkerPool) StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan book.Metadata, resultChannel chan string, exportPath string) {
	for i := 0; i < numberOfWorkers; i++ {
		go y.scrapeWorker(jobChannel, resultChannel, exportPath)
	}
}

func (y *yoinkWorkerPool) scrapeWorker(jobs <-chan book.Metadata, results chan<- string, exportPath string) {
	for metadata := range jobs {
		results <- y.startScraping(metadata, exportPath)
	}
}

func (y *yoinkWorkerPool) startScraping(book book.Metadata, exportPath string) string {
	yoinker := factory.NewYoinker()
	yoinker.StartYoink(book, exportPath)
	return book.Title
}

//BeginMultiConvert scrapes all books with a worker pool of given size and exports them to a given path
func (y *yoinkWorkerPool) BeginMultiConvert(books []book.Metadata, numberOfWorkers int, outputPath string) {
	jobChannel := make(chan book.Metadata, 100)
	resultChannel := make(chan string, 100)

	y.StartScrapeWorkerPool(numberOfWorkers, jobChannel, resultChannel, outputPath)

	for _, metadata := range books {
		jobChannel <- metadata
	}
	close(jobChannel)
	for range books {
		<-resultChannel
	}
	close(resultChannel)
}

//NewScrapeWorkerpool Creates a new instance of a scraper worker pool
func NewScrapeWorkerpool() ScrapeWorkerpool {
	return &yoinkWorkerPool{}
}
