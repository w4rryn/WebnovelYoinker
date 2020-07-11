package terminal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
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
	logErr(err)
}

func scrapeCommand(c *cli.Context) error {
	fmt.Println("Starting conversion.")
	fmt.Println("Status:")
	jobs := getBookConfigs(c.String("in"))
	numOfJobs := len(jobs)
	bar := progressbar.Default(int64(numOfJobs) * 2)
	addBarStep := func(s string) {
		bar.Add(1)
	}
	yoinker.OnScrapeStart = append(yoinker.OnScrapeStart, addBarStep)
	yoinker.OnExportFinished = append(yoinker.OnExportFinished, addBarStep)

	workerPool := yoinker.NewScrapeWorkerpool()
	workerPool.BeginMultiConvert(jobs, c.Int("r"), c.String("out"))
	fmt.Println("Finished")
	return nil
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBookConfigs(path string) []yoinker.BookMetadata {
	books := []yoinker.BookMetadata{}
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
