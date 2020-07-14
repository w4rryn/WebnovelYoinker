package terminal

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// StartTerminal start a terminal applcation to use Webnovel Yoinker
func StartTerminal() {
	app := cli.NewApp()
	app.Name = "WebnovelYoinker terminal"
	app.Usage = "Lets you download webnovels and exports them as epub or pdf"
	app.Commands = []*cli.Command{
		{
			Name:   "scrape",
			Usage:  "Downloads specified books and exports them",
			Flags:  scrapeFlags,
			Action: scrapeCommand,
		},
		// {
		// 	Name:   "yoink",
		// 	Usage:  "Attempts to scrape a single volume from a specified url",
		// 	Flags:  singleScrapeFlags,
		// 	Action: singleScrapeCommand,
		// },
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
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
