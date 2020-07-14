package export

import (
	"fmt"
	"html"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bmaupin/go-epub"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//epubExporter exports a volume a epub
type epubExporter struct {
	epubExport *epub.Epub
}

//Export exports a valume as epub
func (e *epubExporter) Export(metadata book.Metadata, path string, chapters []book.Chapter) string {
	var (
		coverImage string
		waiter     sync.WaitGroup
	)

	e.epubExport = epub.NewEpub(metadata.Title)
	cssPath, err := e.epubExport.AddCSS("https://raw.githubusercontent.com/lethal-bacon0/WebnovelYoinker/master/assets/ebookstyle.css", "stylesheet.css")
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		waiter.Add(1)
		var err error
		coverImage, err = e.epubExport.AddImage(metadata.Cover, "")
		if err != nil {
			log.Fatal(err.Error())
		}
		e.epubExport.SetCover(coverImage, "")
		e.epubExport.SetAuthor(metadata.Author)
		e.epubExport.SetLang(metadata.Language)
		waiter.Done()
	}()

	for i, chapter := range chapters {
		if chapter.ChapterName == "" {
			chapter.ChapterName = fmt.Sprintf("Chapter %v", i+1)
		}
		e.epubExport.AddSection(e.addChapter(chapter), chapter.ChapterName, "", cssPath)
	}

	waiter.Wait()

	exportPath := filepath.Join(path, metadata.Title+".epub")
	if err != nil {
		log.Fatal(err)
	}
	err = e.epubExport.Write(exportPath)
	if err != nil {
	}
	return exportPath
}

func (e *epubExporter) addChapter(chapter book.Chapter) string {
	var parsedContent strings.Builder
	parsedContent.WriteString(fmt.Sprintf("<p><strong> %v </strong></p>", chapter.ChapterName))
	for _, page := range chapter.Content {
		switch page.(type) {
		case *book.PageImage:
			pageImage := page.(*book.PageImage)
			imagePath, err := e.epubExport.AddImage(pageImage.Image, "")
			if err != nil {
				continue
			}
			content := fmt.Sprintf("<div class=\"width\">"+
				"<div class=\"pc\">"+
				"<p>"+
				"<img src=\"%v\" width=\"%v\" height=\"%v\" class=\"calibre1\" alt=\"image\"/>"+
				"</p>"+
				"</div>"+
				"</div>",
				imagePath, pageImage.Width, pageImage.Height)
			parsedContent.WriteString(content)

		case *book.Paragraph:
			par := page.(*book.Paragraph)
			content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
			parsedContent.WriteString(content)
		}
	}
	return parsedContent.String()
}

//NewEpubExporter creates a new epub exporter
func NewEpubExporter() yoinker.IExportStrategy {
	return &epubExporter{}
}
