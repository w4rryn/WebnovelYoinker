package terminal

import (
	"fmt"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/export"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/scrape"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
)

var scrapeFlags = []cli.Flag{
	&cli.PathFlag{
		Name:     "-input-File",
		Value:    "~/",
		Aliases:  []string{"in"},
		Required: true,
	},
	&cli.PathFlag{
		Name:     "-output-directory",
		Value:    "~/Downloads",
		Aliases:  []string{"out"},
		Required: true,
	},

	&cli.IntFlag{
		Name:        "r",
		Usage:       "Sets the number of go routines used to scrape, aka how many should be downloaded simultaenously",
		DefaultText: "3",
		Value:       3,
	},
}

func scrapeCommand(c *cli.Context) error {
	fmt.Println("Starting conversion.")
	fmt.Println("Status:")
	jobs := getBookConfigs(c.String("in"))
	numOfJobs := len(jobs)
	bar := progressbar.Default(int64(numOfJobs))
	jobChannel := make(chan book.Metadata, 100)
	resultChannel := make(chan string, 100)

	for i := 0; i < c.Int("r"); i++ {
		go scrapeWorker(jobChannel, resultChannel, c.Path("out"))
	}

	go func() {
		for _, job := range jobs {
			jobChannel <- job
		}
		close(jobChannel)
	}()

	for i := 0; i < len(jobs); i++ {
		<-resultChannel
		bar.Add(1)
	}
	close(resultChannel)
	fmt.Println("Finished")
	return nil
}

func scrapeWorker(jobs <-chan book.Metadata, results chan<- string, exportPath string) {
	for job := range jobs {
		scraper, err := scrape.GetScraper(job.Website)
		if err != nil {
			fmt.Println(err.Error())
		}
		exporter, err := export.GetExporter(job.Format)
		if err != nil {
			fmt.Println(err.Error())
		}
		yoinkerFactory := yoinker.NewYoinkerFactory(scraper, exporter)
		yoinker := yoinkerFactory.GetYoinker()
		results <- yoinker.StartYoink(job, exportPath)
	}
}
