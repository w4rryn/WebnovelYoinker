package yoinker

import (
	"fmt"
)

//WorkerPoolYoinker provieds interface for yoinker worker pattern
type WorkerPoolYoinker interface {
	StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string)
}

//PoolYoinker Implements WorkerPoolYoinker to provide abstraction for yoinker worer pools
type PoolYoinker struct {
}

//StartScrapeWorkerPool start the yoinker worker pool
func (p *PoolYoinker) StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string) {
	for i := 0; i < numberOfWorkers; i++ {
		go p.worker(jobChannel, resultChannel, exportPath)
	}
}

func (p *PoolYoinker) worker(jobs <-chan BookMetadata, results chan<- string, exportPath string) {
	for metadata := range jobs {
		results <- p.startScraping(metadata, exportPath)
	}
}

func (p *PoolYoinker) startScraping(bookMetadata BookMetadata, exportPath string) string {
	yoinker := NewYoinker()
	fmt.Printf("Start scraping %v \n", bookMetadata.Title)
	yoinker.StartYoink(bookMetadata, exportPath)
	return bookMetadata.Title
}
