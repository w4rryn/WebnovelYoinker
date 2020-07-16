package scrape

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// crimsonmagicNovelScraper is a concrete strategy to scrape a novel from crimsonmagic.com
type crimsonmagicNovelScraper struct {
	chapterUrls   []string
	PrintCallback func(s string)
}

//ScrapeChapter Scrapes all chapters
func (c *crimsonmagicNovelScraper) ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter {
	resp, err := http.Get(chapterURL)
	if err != nil {
		return book.Chapter{}
	}
	root, err := html.Parse(resp.Body)
	chapter := book.Chapter{
		ChapterNumber: chapterNumber,
		URL:           chapterURL,
	}
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
			pageImage := book.NewPageImage(int(height), int(width), scrape.Attr(img, "data-src"))
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
			chapter.Content = append(chapter.Content, book.NewParagraph(scrape.Text(par)))
		}
	}
	return chapter
}

//GetAvailableChapters gets all available Volume information from a url
func (c crimsonmagicNovelScraper) GetAvailableChapters(url string) []book.Volume {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	root, err := html.Parse(response.Body)
	entryContent, ok := scrape.Find(root, scrape.ByClass("entry-content"))
	if !ok {
		return nil
	}
	volumeMatcher := func(n *html.Node) bool {
		if n.DataAtom == atom.P && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n, "style") == "text-align: center;"
		}
		return false
	}
	var volumes []book.Volume
	volumeNodes := scrape.FindAll(entryContent, volumeMatcher)
	for i, volumeNode := range volumeNodes {
		var chapters []book.Chapter
		for _, chapterNode := range scrape.FindAll(volumeNode, scrape.ByTag(atom.A)) {
			chapters = append(chapters, book.Chapter{
				ChapterName: scrape.Text(chapterNode),
				URL:         scrape.Attr(chapterNode, "href"),
			})
		}
		currentVolume := book.Volume{
			Chapters: chapters,
			Metadata: book.Metadata{
				Title: fmt.Sprintf("Volume %v", i+1),
			},
		}
		volumes = append(volumes, currentVolume)
	}
	return volumes
}

//NewCrimsonmagicScraper creates a new crimsonmagic scraper
func NewCrimsonmagicScraper() yoinker.IScrapingStrategy {
	return &crimsonmagicNovelScraper{}
}
