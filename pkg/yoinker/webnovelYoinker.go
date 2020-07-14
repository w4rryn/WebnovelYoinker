package yoinker

import (
	"sync"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//IYoinkerManager Provides Functionality to yoink Webnovels and Webtoons
type IYoinkerManager interface {
	StartYoink(metadata book.Metadata, exportPath string) string
	GetAvailableVolumes(url string, website string) []book.Volume
}

//WebnovelYoinker scrapes webnovels and webtoons and exports them as epub or pdf
type webnovelYoinker struct {
	scraper  IScrapingStrategy
	exporter IExportStrategy
}

// StartYoink start yoinking the specified book
func (w *webnovelYoinker) StartYoink(metadata book.Metadata, path string) string {
	var (
		scrapedChapters []book.Chapter
		waiter          sync.WaitGroup
	)

	jobs := make(chan book.Chapter, 100)
	results := make(chan book.Chapter, 100)
	for i := 0; i < 16; i++ {
		go w.chapterScraperWorker(jobs, results)
	}
	waiter.Add(1)
	go func() {
		for i, chapterURL := range metadata.ChapterURLs {
			jobs <- book.Chapter{ChapterNumber: i, URL: chapterURL}
		}
		close(jobs)
		waiter.Done()
	}()
	for i := 0; i < len(metadata.ChapterURLs); i++ {
		scrapedChapters = append(scrapedChapters, <-results)
	}
	close(results)
	return w.exporter.Export(metadata, path, w.sortChapters(scrapedChapters))
}

func (w *webnovelYoinker) chapterScraperWorker(jobs <-chan book.Chapter, results chan<- book.Chapter) {
	for job := range jobs {
		results <- w.scraper.ScrapeChapter(job.URL, job.ChapterNumber)
	}
}

//GetAvailableVolumes get all available volumes of provided url
func (w *webnovelYoinker) GetAvailableVolumes(url string, website string) []book.Volume {
	return w.scraper.GetAvailableChapters(url)
}

// simple bubble sort for chapter sorting
func (w *webnovelYoinker) sortChapters(chapters []book.Chapter) []book.Chapter {
	var (
		n      = len(chapters)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if chapters[i].ChapterNumber > chapters[i+1].ChapterNumber {
				chapters[i+1], chapters[i] = chapters[i], chapters[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return chapters
}
