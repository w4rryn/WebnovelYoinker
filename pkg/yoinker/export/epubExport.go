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
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/events"
)

//EpubExporter exports a volume a epub
type EpubExporter struct {
	epubExport *epub.Epub
}

//Export exports a valume as epub
func (e *EpubExporter) Export(metadata book.Metadata, path string, chapterChannel <-chan book.Chapter) string {
	go func() {
		events.OnExportStartEvent.Invoke(&yoinker.CtxYoink{
			Volume: book.Volume{
				Metadata: metadata,
			},
		})
	}()
	e.epubExport = epub.NewEpub(metadata.Title)
	cssPath, err := e.epubExport.AddCSS("https://raw.githubusercontent.com/lethal-bacon0/WebnovelYoinker/master/assets/ebookstyle.css", "stylesheet.css")
	if err != nil {
		// invokeError(err)
	}
	var coverImage string
	var waiter sync.WaitGroup
	go func() {
		waiter.Add(1)
		var err error
		coverImage, err = e.epubExport.AddImage(metadata.Cover, "")
		if err != nil {
			// invokeError(err)
		}
		waiter.Done()
	}()

	i := 0
	for chapter := range chapterChannel {
		i++
		if chapter.ChapterName == "" {
			chapter.ChapterName = fmt.Sprintf("Chapter %v", i)
		}
		e.epubExport.AddSection(e.addChapter(chapter), chapter.ChapterName, "", cssPath)
	}

	waiter.Wait()
	e.epubExport.SetCover(coverImage, "")
	e.epubExport.SetAuthor(metadata.Author)
	e.epubExport.SetLang(metadata.Language)
	exportPath := filepath.Join(path, metadata.Title+".epub")
	if err != nil {
		log.Fatal(err)
	}
	err = e.epubExport.Write(exportPath)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		events.OnExportFinishedEvent.Invoke(&yoinker.CtxYoink{
			Volume: book.Volume{
				Metadata: metadata,
			},
		})
	}()
	return exportPath
}

func (e *EpubExporter) addChapter(chapter book.Chapter) string {
	var parsedContent strings.Builder
	parsedContent.WriteString(fmt.Sprintf("<p><strong> %v </strong></p>", chapter.ChapterName))
	for _, page := range chapter.Content {
		switch page.(type) {
		case *book.PageImage:
			pageImage := page.(*book.PageImage)
			imagePath, err := e.epubExport.AddImage(pageImage.Image, "")
			if err != nil {
				// e.makeCallback(err.Error())
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
