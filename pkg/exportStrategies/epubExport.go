package exportstrategies

import (
	"fmt"
	"html"
	"strings"
	"sync"

	"github.com/bmaupin/go-epub"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

//EpubExporter exports a volume a epub
type EpubExporter struct {
	Callback   func(s string)
	epubExport *epub.Epub
}

//Export exports a valume as epub
func (e *EpubExporter) Export(metadata yoinker.BookMetadata, path string, chapterChannel <-chan yoinker.Chapter) string {
	e.epubExport = epub.NewEpub(metadata.Title)
	var coverImage string
	var waiter sync.WaitGroup
	go func() {
		waiter.Add(1)
		var err error
		coverImage, err = e.epubExport.AddImage(metadata.Cover, "")
		if err != nil {
			panic(err)
		}
		waiter.Done()
	}()

	i := 0
	for chapter := range chapterChannel {
		i++
		e.epubExport.AddSection(e.addChapter(chapter), fmt.Sprintf("Chapter %d", i), "", "")
	}

	waiter.Wait()
	e.epubExport.SetCover(coverImage, "")
	e.epubExport.SetAuthor(metadata.Author)
	e.epubExport.SetLang(metadata.Language)
	exportPath := metadata.Title + ".epub"
	err := e.epubExport.Write(exportPath)
	if err != nil {
		panic(err)
	}
	return exportPath
}

func (e EpubExporter) checkError(err error) {
	if err != nil {
		e.Callback(err.Error())
	}
}

func (e *EpubExporter) addChapter(chapter yoinker.Chapter) string {
	var parsedContent strings.Builder
	for _, paragraph := range chapter.Content {
		switch paragraph.(type) {
		case *yoinker.PageImage:
			pageImage := paragraph.(*yoinker.PageImage)
			imagePath, err := e.epubExport.AddImage(pageImage.Image, "")
			e.checkError(err)
			content := fmt.Sprintf("<p style=\"page-break-before: always\"><img src=\"%v\" width=\"%v\" height=\"%v\"/></p>", imagePath, pageImage.Width, pageImage.Height)
			parsedContent.WriteString(content)

		case *yoinker.Paragraph:
			par := paragraph.(*yoinker.Paragraph)
			content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
			parsedContent.WriteString(content)
		}
	}
	return parsedContent.String()
}
