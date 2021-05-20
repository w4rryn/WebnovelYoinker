package scrape

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type FullyBookedTLSScraper struct {
}

//ScrapeChapter scrapes all chapters
func (f *FullyBookedTLSScraper) ScrapeChapter(chapterURL string, chapterNumber int) book.Chapter {
	resp, err := http.Get(chapterURL)
	if err != nil {
		return book.Chapter{}
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	chapter := book.Chapter{
		ChapterNumber: chapterNumber,
		URL:           chapterURL,
	}

	matcher := func(node *html.Node) bool {
		a := node.DataAtom
		return a == atom.P || a == atom.Figure || a == atom.Hr
	}

	imgfilter := scrape.ByTag(atom.Figure)
	main, _ := scrape.Find(root, scrape.ByClass("entry-content"))
	imageNumber := 0
	foundHr := false
	chapterNameFound := false
	for _, p := range scrape.FindAll(main, matcher) {
		if p.DataAtom == atom.Hr {
			foundHr = true
			continue
		}

		if p.DataAtom == atom.P && !foundHr {
			continue
		}

		if img, ok := scrape.Find(p, imgfilter); ok {
			sizes := scrape.Attr(img, "data-orig-size")
                        if sizes == "" {
				sizes = "1127,1600"
			}
			size := strings.Split(sizes, ",")
			width, err := strconv.ParseInt(size[0], 0, 64)
			if err != nil {
				width = 1127
			}
			height, err := strconv.ParseInt(size[1], 0, 64)
			if err != nil {
				width = 1600
			}
			page := book.NewPageImage(int(height), int(width), scrape.Attr(img, "src"))
			chapter.Images = append(chapter.Images, page.Image)
			chapter.Content = append(chapter.Content, page)
			imageNumber++
		} else {
			strong := scrape.ByTag(atom.Strong)
			if chapterName, ok := scrape.Find(p, strong); ok && !chapterNameFound {
				chapter.ChapterName = scrape.Text(chapterName)
				chapterNameFound = true
			} else {
				par := book.NewParagraph(scrape.Text(p))
				chapter.Content = append(chapter.Content, par)
			}
		}
	}

	return chapter
}

func (f *FullyBookedTLSScraper) GetAvailableChapters(url string) []book.Volume {
	panic("Not implemented") // TODO: implement  this
}

func NewFullyBookedTLSSCraper() *FullyBookedTLSScraper {
	return &FullyBookedTLSScraper{}
}
