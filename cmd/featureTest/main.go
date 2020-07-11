// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
// 	"gopkg.in/yaml.v2"
// )

// func main() {
// 	jobChannel := make(chan yoinker.BookMetadata, 100)
// 	resultChannel := make(chan string, 100)

// 	jobs := getBookConfigs()
// 	workerPool := yoinker.NewScrapeWorkerpool()
// 	workerPool.StartScrapeWorkerPool(3, jobChannel, resultChannel, "~/Downloads/books")

// 	for _, metadata := range jobs {
// 		jobChannel <- metadata
// 	}
// 	close(jobChannel)
// 	for range jobs {
// 		result := <-resultChannel
// 		fmt.Printf("Finished scraping %v\n", result)
// 	}
// 	close(resultChannel)
// }

// func getBookConfigs() []yoinker.BookMetadata {
// 	books := []yoinker.BookMetadata{}
// 	rawBooks, err := ioutil.ReadFile("exports.yaml")
// 	if err != nil {
// 		log.Fatalf("%v: %v", err.Error(), err)
// 	}

// 	err = yaml.Unmarshal(rawBooks, &books)
// 	if err != nil {
// 		log.Fatalf("Unknown format: %v", err)
// 	}

// 	return books
// }
