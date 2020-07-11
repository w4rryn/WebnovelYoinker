package yoinker

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// crimsonmagicNovelScraper is a concrete strategy to scrape a novel from cromsonmagic.com
type crimsonmagicNovelScraper struct {
	chapterUrls   []string
	PrintCallback func(s string)
}

//BeginScrape Scrapes all chapters
func (c *crimsonmagicNovelScraper) BeginScrape(chapterURLs []string, chapterChannel chan<- chapter) {
	chapterStrings := strings.Join(chapterURLs, ",")
	invokeYoinkerScrapeEvent(OnScrapeStart, chapterStrings)
	for _, chapterURL := range chapterURLs {
		resp, err := http.Get(chapterURL)
		if err != nil {
			invokeError(err)
		}
		root, err := html.Parse(resp.Body)
		if err != nil {
			invokeError(err)
		}
		chapter := c.getChapter(root)
		chapterChannel <- chapter
	}

	invokeYoinkerScrapeEvent(OnChapterScraped, chapterStrings)
	close(chapterChannel)
}

func (c crimsonmagicNovelScraper) getChapter(root *html.Node) chapter {
	var chapter chapter
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
			pageImage := newPageImage(int(height), int(width), scrape.Attr(img, "data-src"))
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
			chapter.Content = append(chapter.Content, newParagraph(scrape.Text(par)))
		}
	}
	return chapter
}
