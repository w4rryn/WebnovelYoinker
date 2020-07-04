package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/bmaupin/go-epub"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	chapters := []string{
		"https://www.crimsonmagic.me/joshikousei/jk-1-p/",
		"https://www.crimsonmagic.me/joshikousei/JK-1-1/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-2/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-3/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-4/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-5/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-6/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-7/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-8/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-9/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-10/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-11/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-12/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-13/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-14/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-15/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-16/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-17/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-e/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-a/",
		"https://www.crimsonmagic.me/joshikousei/jk-1-ss/",
	}

	// chapters := []string{
	// 	"https://www.crimsonmagic.me/archive/gifting-1-p/",
	// 	"https://www.crimsonmagic.me/archive/gifting-1-1/",
	// 	"https://www.crimsonmagic.me/archive/gifting-1-2/",
	// 	"https://www.crimsonmagic.me/archive/gifting-1-3/",
	// 	"https://www.crimsonmagic.me/archive/gifting-1-e/",
	// }

	volume, err := scrapeVolume(chapters, func(s string) {
		fmt.Println(s)
	})
	check(err)
	exportEpub(volume)
}

func scrapeVolume(chapterURLs []string, callback func(s string)) (*Volume, error) {
	volume := Volume{
		Chapters: getChapters(chapterURLs, callback),
		Author:   "yuNS",
		Cover:    "https://crimsonmagicme.files.wordpress.com/2018/08/cover1.jpg",
		Language: "English",
		Title:    "I Shaved. Then I Brought a High School Girl Home. Volume 1",
		Year:     "2018",
	}
	return &volume, nil
}

//TODO make this recursively
//TODO some images are missing
//TODO start to make this in a proper application
func getChapters(chapterURLs []string, callback func(s string)) []Chapter {
	var chapters []Chapter
	for chapNum, chapterURL := range chapterURLs {
		callback(fmt.Sprintf("Scraping chapter: %v. URL: %v", chapNum, chapterURL))
		resp, err := http.Get(chapterURL)
		check(err)
		root, err := html.Parse(resp.Body)
		check(err)

		matcherClass := scrape.ByClass("entry-content")
		matcherParagraph := scrape.ByTag(atom.P)
		class, _ := scrape.Find(root, matcherClass)
		paragraphNodes := scrape.FindAll(class, matcherParagraph)

		var chapter Chapter
		var chapterContent strings.Builder
		imageNumber := 0

		for _, par := range paragraphNodes {
			imageFilter := scrape.ByTag(atom.Img)
			if img, ok := scrape.Find(par, imageFilter); ok {
				chapter.Images = append(chapter.Images, scrape.Attr(img, "data-src"))
				image := "<p style=\"page-break-before: always\">" +
					"<img class=\"%v\" src=\"{{index .Images %v}}\" width=\"%v\" height=\"%v\"/>" +
					"</p>"
				content := fmt.Sprintf(image, scrape.Attr(img, "class"), imageNumber, scrape.Attr(img, "width"), scrape.Attr(img, "height"))
				chapterContent.WriteString(content)
				imageNumber++
				continue
			}
			chapterContent.WriteString("<p>" + scrape.Text(par) + "</p>")
		}
		chapter.Content = chapterContent.String()
		chapters = append(chapters, chapter)
	}
	return chapters
}

func exportEpub(volume *Volume) {
	// Create a new EPUB
	e := epub.NewEpub(volume.Title)
	coverImage, err := e.AddImage(volume.Cover, "")
	check(err)
	e.SetCover(coverImage, "")
	e.SetAuthor(volume.Author)
	e.SetLang(volume.Language)

	for i, chapter := range volume.Chapters {
		var localImages []string
		for _, imageURL := range chapter.Images {
			epubImage, err := e.AddImage(imageURL, "")
			check(err)
			localImages = append(localImages, epubImage)
		}
		chapter.Images = localImages
		chapterName := fmt.Sprintf("Chapter %d", i+1)
		chapterTemplate := template.New("Template_1")
		chapterTemplate, _ = chapterTemplate.Parse(chapter.Content)
		var stringBuffer bytes.Buffer
		err := chapterTemplate.Execute(&stringBuffer, chapter)
		check(err)

		e.AddSection(stringBuffer.String(), chapterName, "", "")
	}
	err = e.Write(volume.Title)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Volume represents a book that consists of chapters
type Volume struct {
	Cover    string
	Chapters []Chapter
	Author   string
	Year     string
	Language string
	Title    string
}

// Chapter represents a chapter in a book.
// Chapter content must be valid xhtml
type Chapter struct {
	Images  []string
	Content string
}
