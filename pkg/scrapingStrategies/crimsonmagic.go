package scrapingstrategies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// CrimsonmagicNovelScraper is a concrete strategy to scrape a novel from cromsonmagic.com
type CrimsonmagicNovelScraper struct {
	chapterUrls   []string
	PrintCallback func(s string)
}

//BeginScrape Scrapes all chapters
func (c *CrimsonmagicNovelScraper) BeginScrape(chapterURLs []string, chapterChannel chan<- yoinker.Chapter) {
	for _, chapterURL := range chapterURLs {
		c.makeCallback(fmt.Sprintf("Downloading chapter: %v", chapterURL))
		resp, err := http.Get(chapterURL)
		if err != nil {
			c.makeCallback(err.Error())
		}
		root, err := html.Parse(resp.Body)
		if err != nil {
			c.makeCallback(err.Error())
		}
		chapterChannel <- c.getChapter(root)
	}

	close(chapterChannel)
}

func (c CrimsonmagicNovelScraper) getChapter(root *html.Node) yoinker.Chapter {
	var chapter yoinker.Chapter
	mainContentMatcher := scrape.ByClass("entry-content")
	paragraphMatcher := scrape.ByTag(atom.P)
	class, _ := scrape.Find(root, mainContentMatcher)
	imageNumber := 0
	var chapterNameFound bool
	for _, par := range scrape.FindAll(class, paragraphMatcher) {
		imageFilter := scrape.ByTag(atom.Img)
		if img, ok := scrape.Find(par, imageFilter); ok {
			widthAtr := scrape.Attr(img, "width")
			width, err := strconv.ParseInt(widthAtr, 0, 64)
			if err != nil {
				width = 1127
			}
			heightAttr := scrape.Attr(img, "height")
			height, err := strconv.ParseInt(heightAttr, 0, 64)
			if err != nil {
				height = 1600
			}
			pageImage := yoinker.NewPageImage(int(height), int(width), scrape.Attr(img, "data-src"))
			chapter.Images = append(chapter.Images, pageImage.Image)
			chapter.Content = append(chapter.Content, pageImage)
			imageNumber++
			continue
		} else {
			spanMatcher := scrape.ByTag(atom.Span)
			if span, ok := scrape.Find(par, spanMatcher); ok && scrape.Attr(span, "style") == "color:#ffffff;" {
				continue
			}
			strongMatcher := scrape.ByTag(atom.Strong)
			boldMatcher := scrape.ByTag(atom.B)
			if chapterName, ok := scrape.Find(par, strongMatcher); ok && !chapterNameFound {
				chapterNameFound = true
				chapter.ChapterName = scrape.Text(chapterName)
				continue
			} else if chapterName, ok := scrape.Find(par, boldMatcher); ok && !chapterNameFound {
				chapterNameFound = true
				chapter.ChapterName = scrape.Text(chapterName)
				continue
			}
			chapter.Content = append(chapter.Content, yoinker.NewParagraph(scrape.Text(par)))
		}
	}
	return chapter
}

func (c CrimsonmagicNovelScraper) makeCallback(s string) {
	if c.PrintCallback != nil {
		c.PrintCallback(s)
	}
}
