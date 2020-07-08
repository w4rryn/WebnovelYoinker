package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	yc "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/yoinkerCore"
	"gopkg.in/yaml.v2"
)

func main() {
	jobChannel := make(chan yc.BookMetadata, 100)
	resultChannel := make(chan string, 100)

	jobs := getBookConfigs()
	startScrapeWorkerPool(3, jobChannel, resultChannel)

	for _, metadata := range jobs {
		jobChannel <- metadata
	}
	close(jobChannel)
	for range jobs {
		result := <-resultChannel
		fmt.Printf("Finished scraping %v\n", result)
	}
	close(resultChannel)
}

func startScrapeWorkerPool(numberOfWorkers int, jobChannel chan yc.BookMetadata, resultChannel chan string) {
	for i := 0; i < numberOfWorkers; i++ {
		go worker(jobChannel, resultChannel)
	}
}

func worker(jobs <-chan yc.BookMetadata, results chan<- string) {
	for metadata := range jobs {
		results <- startScraping(metadata)
	}
}

func startScraping(bookMetadata yc.BookMetadata) string {
	yoinker := yoinker.NewYoinker()
	fmt.Printf("Start scraping %v \n", bookMetadata.Title)
	yoinker.StartYoink(bookMetadata)
	return bookMetadata.Title
}

func getBookConfigs() []yc.BookMetadata {
	books := []yc.BookMetadata{}
	rawBooks, err := ioutil.ReadFile("exports.yaml")
	if err != nil {
		log.Fatalf("%v: %v", err.Error(), err)
	}

	err = yaml.Unmarshal(rawBooks, &books)
	if err != nil {
		log.Fatalf("Unknown format: %v", err)
	}

	return books
}
