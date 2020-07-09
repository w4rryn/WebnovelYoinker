package yoinker

import (
	"fmt"
	"html"
	"strings"
	"sync"

	"github.com/bmaupin/go-epub"
)

//epubExporter exports a volume a epub
type epubExporter struct {
	epubExport *epub.Epub
	// PrintCallback func(s string)
}

//Export exports a valume as epub
func (e *epubExporter) Export(metadata BookMetadata, path string, chapterChannel <-chan chapter) string {
	e.epubExport = epub.NewEpub(metadata.Title)
	cssPath, err := e.epubExport.AddCSS("../assets/ebookstyle.css", "stylesheet.css")
	if err != nil {
		// e.PrintCallback(err.Error())
	}
	var coverImage string
	var waiter sync.WaitGroup
	go func() {
		waiter.Add(1)
		var err error
		coverImage, err = e.epubExport.AddImage(metadata.Cover, "")
		if err != nil {
			// e.makeCallback(err.Error())
		}
		waiter.Done()
	}()

	i := 0
	for chapter := range chapterChannel {
		i++
		if chapter.ChapterName == "" {
			chapter.ChapterName = fmt.Sprintf("Chapter %v", i)
		}
		// e.makeCallback(fmt.Sprintf("Adding chapter %v of %v.", chapter.ChapterName, metadata.Title))
		e.epubExport.AddSection(e.addChapter(chapter), chapter.ChapterName, "", cssPath)
	}

	waiter.Wait()
	e.epubExport.SetCover(coverImage, "")
	e.epubExport.SetAuthor(metadata.Author)
	e.epubExport.SetLang(metadata.Language)
	// e.PrintCallback(fmt.Sprintf("Exporting %v", metadata.Title))
	exportPath := metadata.Title + ".epub"
	err = e.epubExport.Write(exportPath)
	if err != nil {
		// e.makeCallback(err.Error())
	}
	// e.PrintCallback(fmt.Sprintf("Finished exporting %v", metadata.Title))
	return exportPath
}

func (e *epubExporter) addChapter(chapter chapter) string {
	var parsedContent strings.Builder
	parsedContent.WriteString(fmt.Sprintf("<p><strong> %v </strong></p>", chapter.ChapterName))
	for _, page := range chapter.Content {
		switch page.(type) {
		case *pageImage:
			pageImage := page.(*pageImage)
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

		case *paragraph:
			par := page.(*paragraph)
			content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
			parsedContent.WriteString(content)
		}
	}
	return parsedContent.String()
}
