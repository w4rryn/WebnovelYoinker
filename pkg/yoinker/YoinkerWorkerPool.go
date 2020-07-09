package yoinker

import (
	"fmt"
)

//WorkerPoolYoinker provieds interface for yoinker worker pattern
type WorkerPoolYoinker interface {
	StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string)
}

//PoolYoinker Implements WorkerPoolYoinker to provide abstraction for yoinker worer pools
type PoolYoinker struct {
}

//StartScrapeWorkerPool start the yoinker worker pool
func (p *PoolYoinker) StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string) {
	for i := 0; i < numberOfWorkers; i++ {
		go p.worker(jobChannel, resultChannel)
	}
}

func (p *PoolYoinker) worker(jobs <-chan BookMetadata, results chan<- string) {
	for metadata := range jobs {
		results <- p.startScraping(metadata)
	}
}

func (p *PoolYoinker) startScraping(bookMetadata BookMetadata) string {
	yoinker := NewYoinker()
	fmt.Printf("Start scraping %v \n", bookMetadata.Title)
	yoinker.StartYoink(bookMetadata)
	return bookMetadata.Title
}
