package terminal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/events"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/worker"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// StartTerminal start a terminal applcation to use Webnovel Yoinker
func StartTerminal() {
	app := cli.NewApp()
	app.Name = "WebnovelYoinker terminal"
	app.Usage = "Lets you download webnovels and exports them as epub or pdf"

	scrapeFlags := []cli.Flag{
		&cli.PathFlag{
			Name:     "in",
			Value:    "~/",
			Required: true,
		},
		&cli.PathFlag{
			Name:     "out",
			Value:    "~/Downloads",
			Required: true,
		},

		&cli.IntFlag{
			Name:        "r",
			Usage:       "Sets the number of go routines used to scrape, aka how many should be downloaded simultaenously",
			DefaultText: "3",
			Value:       3,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "scrape",
			Usage:  "Downloads specified books and exports them",
			Flags:  scrapeFlags,
			Action: scrapeCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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

func logErr(y *yoinker.CtxYoink) {
	if y.Error != nil {
		fmt.Println(y.Error.Error())
	}
}

func getBookConfigs(path string) []book.Metadata {
	books := []book.Metadata{}
	rawBooks, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%v: %v", err.Error(), err)
	}

	err = yaml.Unmarshal(rawBooks, &books)
	if err != nil {
		log.Fatalf("Unknown format: %v", err)
	}

	return books
}
