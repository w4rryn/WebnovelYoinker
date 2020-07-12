package terminal

import (
	"fmt"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/events"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/worker"
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

	events.OnErrorEvent.Add(logErr)
	fmt.Println("Starting conversion.")
	fmt.Println("Status:")
	jobs := getBookConfigs(c.String("in"))
	numOfJobs := len(jobs)
	bar := progressbar.Default(int64(numOfJobs) * 2)
	addBarStep := func(c *yoinker.CtxYoink) {
		bar.Add(1)
	}
	events.OnExportFinishedEvent.Add(addBarStep)
	events.OnVolumeScrapedEvent.Add(addBarStep)

	workerPool := worker.NewScrapeWorkerpool()
	workerPool.BeginMultiConvert(jobs, c.Int("r"), c.String("out"))
	fmt.Println("Finished")
	return nil
}
