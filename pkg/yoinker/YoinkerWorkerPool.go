package yoinker

//ScrapeWorkerpool privides functionality to create a background worker to export multiple books
type ScrapeWorkerpool interface {
	StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string)
	BeginMultiConvert(books []BookMetadata, numberOfWorkers int, outputPath string)
}

type yoinkerWorkerpoolEvent func(s string)

//YoinkWorkerPool workerpool to scrape and export books
type YoinkWorkerPool struct {
	ExportFinishedEvent   yoinkerWorkerpoolEvent
	ExportStartEvent      yoinkerWorkerpoolEvent
	ScrapeStartEvent      yoinkerWorkerpoolEvent
	ScrapeFinishedEvent   yoinkerWorkerpoolEvent
	ChapterScrapedEvent   yoinkerWorkerpoolEvent
	ParagraphScrapedEvent yoinkerWorkerpoolEvent
}

//StartScrapeWorkerPool start the yoinker worker pool
func (y *YoinkWorkerPool) StartScrapeWorkerPool(numberOfWorkers int, jobChannel chan BookMetadata, resultChannel chan string, exportPath string) {
	for i := 0; i < numberOfWorkers; i++ {
		go y.scrapeWorker(jobChannel, resultChannel, exportPath)
	}
}

func (y *YoinkWorkerPool) scrapeWorker(jobs <-chan BookMetadata, results chan<- string, exportPath string) {
	for metadata := range jobs {
		results <- y.startScraping(metadata, exportPath)
	}
}

func (y *YoinkWorkerPool) startScraping(bookMetadata BookMetadata, exportPath string) string {
	yoinker := NewYoinker()
	y.invokeEvent(y.ScrapeStartEvent, bookMetadata.Title)
	yoinker.StartYoink(bookMetadata, exportPath)
	return bookMetadata.Title
}

//BeginMultiConvert scrapes all books with a worker pool of given size and exports them to a given path
func (y *YoinkWorkerPool) BeginMultiConvert(books []BookMetadata, numberOfWorkers int, outputPath string) {
	jobChannel := make(chan BookMetadata, 100)
	resultChannel := make(chan string, 100)

	y.StartScrapeWorkerPool(numberOfWorkers, jobChannel, resultChannel, outputPath)

	for _, metadata := range books {
		jobChannel <- metadata
	}
	close(jobChannel)
	for range books {
		result := <-resultChannel
		y.invokeEvent(y.ScrapeFinishedEvent, result)
	}
	close(resultChannel)
}

func (y *YoinkWorkerPool) invokeEvent(event yoinkerWorkerpoolEvent, parameter string) {
	if event != nil {
		event(parameter)
	}
}
