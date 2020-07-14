package scrape

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

func sortChapters(chapters []book.Chapter) []book.Chapter {
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
