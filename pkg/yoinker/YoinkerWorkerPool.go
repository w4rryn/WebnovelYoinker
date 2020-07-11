package yoinker

import (
	"fmt"
)

//StartScrapeWorkerPool start the yoinker worker pool
func StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string) {
	for i := 0; i < numberOfWorkers; i++ {
		go scrapeWorker(jobChannel, resultChannel, exportPath)
	}
}

func scrapeWorker(jobs <-chan BookMetadata, results chan<- string, exportPath string) {
	for metadata := range jobs {
		results <- startScraping(metadata, exportPath)
	}
}

func startScraping(bookMetadata BookMetadata, exportPath string) string {
	yoinker := NewYoinker()
	consoleLog(fmt.Sprintf("Start scraping %v \n", bookMetadata.Title))
	yoinker.StartYoink(bookMetadata, exportPath)
	return bookMetadata.Title
}

//BeginMultiConvert scrapes all books with a worker pool of given size and exports them to a given path
func BeginMultiConvert(books []BookMetadata, numberOfWorkers int, outputPath string) {
	jobChannel := make(chan BookMetadata, 100)
	resultChannel := make(chan string, 100)

	StartScrapeWorkerPool(numberOfWorkers, jobChannel, resultChannel, outputPath)

	for _, metadata := range books {
		jobChannel <- metadata
	}
	close(jobChannel)
	for range books {
		result := <-resultChannel
		consoleLog(fmt.Sprintf("Finished scraping %v\n", result))
	}
	close(resultChannel)
}
