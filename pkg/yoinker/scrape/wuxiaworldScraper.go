package scrape

import (
	"fmt"
	"net/http"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//wuxiaWorldScraper scraping strategy for Wuxiaworld.com
type wuxiaWorldScraper struct {
}

// BeginScrape(metadata BookMetadata, chapterURLs []string) (*Volume, error)
func (w *wuxiaWorldScraper) ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter {
	var (
		chapterName         string
		chapterContent      []book.PageObject
		chapterImages       []string
		chapterTitleMatcher = func(n *html.Node) bool {
			if n.DataAtom == atom.H4 && n.Parent != nil && n.Parent.Parent != nil {
				return scrape.Attr(n, "class") == "caption clearfix"
			}
			return false
		}
	)
	response, err := http.Get(chapterURL)
	if err != nil {
		return book.Chapter{}
	}
	root, err := html.Parse(response.Body)
	if chapterNameNode, ok := scrape.Find(root, chapterTitleMatcher); ok {
		chapterName = scrape.Text(chapterNameNode)
	} else {
		chapterName = fmt.Sprintf("Chapter %v", chapterNumber)
	}
	if contentNode, ok := scrape.Find(root, scrape.ById("chapter-content")); ok {
		for i, paragraphNode := range scrape.FindAll(contentNode, scrape.ByTag(atom.P)) {
			chapterContent = append(chapterContent, &book.Paragraph{
				ID:      uint(i),
				Content: scrape.Text(paragraphNode),
			})
		}
	}

	return book.Chapter{
		ChapterNumber: chapterNumber,
		ChapterName:   chapterName,
		URL:           chapterURL,
		Content:       chapterContent,
		Images:        chapterImages,
	}
}

//GetAwvailableChapters returns an array with all possible chapters
func (w *wuxiaWorldScraper) GetAvailableChapters(url string) []book.Volume {
	panic("Not implemented") // TODO: implement  this
}

//NewWuxiaScraper creates a new wuxia scraper strategy
func NewWuxiaScraper() yoinker.IScrapingStrategy {
	return &wuxiaWorldScraper{}
}
