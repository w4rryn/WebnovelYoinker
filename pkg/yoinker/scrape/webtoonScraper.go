package scrape

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//WebtoonScraper scrapes webtoons from webtoons.com
type WebtoonScraper struct {
}

// ScrapeChapter scrapes given chapter from webtoons.com
func (w WebtoonScraper) ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter {

	var (
		chapterName    string
		chapterContent []book.PageObject
		chapterImages  []string
		contentMatcher = func(n *html.Node) bool {
			if n.DataAtom == atom.Img && n.Parent != nil && n.Parent.Parent != nil && n.Parent.Parent.Parent != nil && n.Parent.Parent.Parent.Parent != nil {
				return scrape.Attr(n.Parent.Parent.Parent.Parent, "id") == "content"
			}
			return false
		}
		chapterTitleMatcher = func(n *html.Node) bool {
			if n.DataAtom == atom.H1 && n.Parent != nil && n.Parent.Parent != nil {
				return scrape.Attr(n, "class") == "subj_episode"
			}
			return false
		}
		// chapter = book.Chapter{
		// 	ChapterNumber: chapterNumber,
		// 	URL:           chapterURL,
		// }
	)

	respnse, err := http.Get(chapterURL)
	if err != nil {
		return book.Chapter{}
	}
	mainNode, err := html.Parse(respnse.Body)
	if err != nil {
		return book.Chapter{}
	}
	if chapterNameNode, ok := scrape.Find(mainNode, chapterTitleMatcher); ok {
		chapterName = scrape.Text(chapterNameNode)
	} else {
		chapterName = fmt.Sprintf("Chapter %v", chapterNumber)
	}
	for _, panelNode := range scrape.FindAll(mainNode, contentMatcher) {
		widthAtr := scrape.Attr(panelNode, "width")
		width, err := strconv.ParseInt(widthAtr, 0, 64)
		if err != nil {
			width = 1127
		}
		heightAttr := scrape.Attr(panelNode, "height")
		height, err := strconv.ParseInt(heightAttr, 0, 64)
		if err != nil {
			height = 1600
		}
		url := scrape.Attr(panelNode, "data-url")
		//TODO: save image from url to cache and use cache path in pageImage.
		pageImage := book.NewPageImage(int(height), int(width), url)
		chapterImages = append(chapterImages, pageImage.Image)
		chapterContent = append(chapterContent, pageImage)
	}

	return book.Chapter{
		ChapterNumber: chapterNumber,
		ChapterName:   chapterName,
		URL:           chapterURL,
		Content:       chapterContent,
		Images:        chapterImages,
	}
}

//GetAvailableChapters returns an array with all possible chapters
func (w WebtoonScraper) GetAvailableChapters(url string) []book.Volume {
	panic("not implemented") // TODO: Implement
}

//NewWebtoonScraper creates a new instance of a webtoon scraper
func NewWebtoonScraper() *WebtoonScraper {
	return &WebtoonScraper{}
}
