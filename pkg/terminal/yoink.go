package terminal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/factories"
	"github.com/urfave/cli/v2"
)

var singleScrapeFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "url",
		Aliases:  []string{"u"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     "website",
		Aliases:  []string{"w"},
		Required: true,
	},
}

//TODO: this needs a metric ton of input sanitization
func singleScrapeCommand(c *cli.Context) error {
	var (
		scraper        yoinker.IScrapingStrategy
		volumes        []book.Volume
		chapterUrls    []string
		metadata       book.Metadata
		yoinkerFactory yoinker.IYoinkerFactory
		yoink          yoinker.IYoinkerManager
	)
	scraper, err := factories.GetScraper(c.String("website"))
	if err != nil {
		return err
	}
	volumes = scraper.GetAvailableChapters(c.String("url"))
	n := 0
	for _, volume := range volumes {
		fmt.Println(volume.Metadata.Title)
		for _, chapter := range volume.Chapters {
			n++
			chapterUrls = append(chapterUrls, chapter.URL)
			fmt.Printf("%v %v: %v\n", n, chapter.ChapterName, chapter.URL)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	lowerBound, upperBound, err := getChapterInterval(reader)
	if err != nil {
		return err
	}

	metadata = *getMetadata(reader)
	metadata.ChapterURLs = getSelectedUrls(lowerBound, upperBound, chapterUrls)
	exporter, err := getExportFormat(reader)
	if err != nil {
		return err
	}
	yoinkerFactory = yoinker.NewYoinkerFactory(scraper, *exporter)
	yoink = yoinkerFactory.GetYoinker()
	yoink.StartYoink(metadata, getExportPath(reader))

	return nil
}

func getChapterInterval(reader *bufio.Reader) (int, int, error) {
	fmt.Print("Download chapters (Format: 1-n;): \n")
	chapterIntervalRaw, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	chapterInterval := strings.Split(strings.TrimSpace(chapterIntervalRaw), "-")
	lower, err := strconv.Atoi(chapterInterval[0])
	if err != nil {
		return 0, 0, err
	}
	upper, err := strconv.Atoi(chapterInterval[1])
	if err != nil {
		return 0, 0, err
	}
	return lower, upper, nil
}

func getSelectedUrls(lower, upper int, allChapters []string) []string {
	if lower < 1 {
		panic("I am too lazy to handle this error properly so fuck off")
	}
	return allChapters[lower-1 : upper]
}

func getExportFormat(reader *bufio.Reader) (*yoinker.IExportStrategy, error) {
	var formats = map[string]string{
		"1": "epub",
		"2": "pdf",
	}
	fmt.Println("Export format:")
	fmt.Println("1 epub")
	fmt.Println("2 pdf")
	formatRaw, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	exporter, err := factories.GetExporter(formats[strings.TrimSpace(formatRaw)])
	if err != nil {
		return nil, err
	}
	return &exporter, nil
}

func getMetadata(reader *bufio.Reader) *book.Metadata {
	fmt.Println("Title:")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Author:")
	author, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Language:")
	language, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Year:")
	year, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Cover URL:")
	coverURL, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return nil
	}

	return &book.Metadata{
		Title:    strings.TrimSpace(title),
		Author:   strings.TrimSpace(author),
		Language: strings.TrimSpace(language),
		Year:     strings.TrimSpace(year),
		Cover:    strings.TrimSpace(coverURL),
	}
}

func getExportPath(reader *bufio.Reader) string {
	fmt.Println("Export Path")
	path, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}

	return strings.TrimSpace(path)
}
