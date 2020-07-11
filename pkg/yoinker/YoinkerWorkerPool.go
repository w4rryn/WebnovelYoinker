package yoinker

//ScrapeWorkerpool privides functionality to create a background worker to export multiple books
type ScrapeWorkerpool interface {
	StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string)
	BeginMultiConvert(books []BookMetadata, numberOfWorkers int, outputPath string)
}

//YoinkWorkerPool workerpool to scrape and export books
type yoinkWorkerPool struct {
}

//StartScrapeWorkerPool start the yoinker worker pool
func (y *yoinkWorkerPool) StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string) {
	for i := 0; i < numberOfWorkers; i++ {
		go y.scrapeWorker(jobChannel, resultChannel, exportPath)
	}
}

func (y *yoinkWorkerPool) scrapeWorker(jobs <-chan BookMetadata, results chan<- string, exportPath string) {
	for metadata := range jobs {
		results <- y.startScraping(metadata, exportPath)
	}
}

func (y *yoinkWorkerPool) startScraping(bookMetadata BookMetadata, exportPath string) string {
	yoinker := NewYoinker()
	yoinker.StartYoink(bookMetadata, exportPath)
	return bookMetadata.Title
}

//BeginMultiConvert scrapes all books with a worker pool of given size and exports them to a given path
func (y *yoinkWorkerPool) BeginMultiConvert(books []BookMetadata, numberOfWorkers int, outputPath string) {
	jobChannel := make(chan BookMetadata, 100)
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
