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
	Callback    func(s string)
	volume      yoinker.Volume
	chapterUrls []string
}

//SetMetadata sets the metadata of the volume
func (c *CrimsonmagicNovelScraper) SetMetadata(author, coverURL, language, title, year string) {
	c.volume = yoinker.Volume{
		Author:   author,
		Cover:    coverURL,
		Language: language,
		Title:    title,
		Year:     year,
	}
}

//BeginScrape Scrapes all chapters
func (c *CrimsonmagicNovelScraper) BeginScrape(chapterURLs []string) (*yoinker.Volume, error) {

	c.chapterUrls = chapterURLs
	chapters, err := c.getChapters()
	c.volume.Chapters = chapters
	if err != nil {
		return nil, err
	}
	return &c.volume, nil
}

//TODO some images are missing
func (c *CrimsonmagicNovelScraper) getChapters() ([]yoinker.Chapter, error) {
	var chapters []yoinker.Chapter
	for chapNum, chapterURL := range c.chapterUrls {
		if c.Callback != nil {
			c.Callback(fmt.Sprintf("Scraping chapter: %v. URL: %v", chapNum, chapterURL))
		}

		resp, err := http.Get(chapterURL)
		if err != nil {
			return nil, err
		}
		root, err := html.Parse(resp.Body)
		if err != nil {
			return nil, err
		}
		chapter := c.getChapter(root)
		chapters = append(chapters, chapter)
	}
	return chapters, nil
}

func (c CrimsonmagicNovelScraper) getChapter(root *html.Node) yoinker.Chapter {
	var chapter yoinker.Chapter
	mainContentMatcher := scrape.ByClass("entry-content")
	paragraphMatcher := scrape.ByTag(atom.P)
	class, _ := scrape.Find(root, mainContentMatcher)
	imageNumber := 0

	for _, par := range scrape.FindAll(class, paragraphMatcher) {
		imageFilter := scrape.ByTag(atom.Img)
		if img, ok := scrape.Find(par, imageFilter); ok {
			width, err := strconv.ParseInt(scrape.Attr(img, "width"), 0, 64)
			if err != nil {
				c.Callback(err.Error())
			}
			height, err := strconv.ParseInt(scrape.Attr(img, "height"), 0, 64)
			if err != nil {
				c.Callback(err.Error())
			}
			pageImage := yoinker.NewPageImage(int(width), int(height), scrape.Attr(img, "data-src"))
			chapter.Images = append(chapter.Images, pageImage.Image)
			chapter.Content = append(chapter.Content, pageImage)
			imageNumber++
			continue
		} else {
			chapter.Content = append(chapter.Content, yoinker.NewParagraph(scrape.Text(par)))
		}
	}
	return chapter
}

func (c CrimsonmagicNovelScraper) callback(s string) {

}
