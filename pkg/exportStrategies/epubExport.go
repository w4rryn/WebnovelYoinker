package exportstrategies

import (
	"fmt"
	"html"
	"strings"

	"github.com/bmaupin/go-epub"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

//EpubExporter exports a volume a epub
type EpubExporter struct {
	Callback   func(s string)
	epubExport *epub.Epub
}

//Export exports a valume as epub
func (e *EpubExporter) Export(volume *yoinker.Volume) error {
	e.epubExport = epub.NewEpub(volume.Metadata.Title)
	coverImage, err := e.epubExport.AddImage(volume.Metadata.Cover, "")
	if err != nil {
		return err
	}
	e.epubExport.SetCover(coverImage, "")
	e.epubExport.SetAuthor(volume.Metadata.Author)
	e.epubExport.SetLang(volume.Metadata.Language)

	for i, chapter := range volume.Chapters {
		e.AddChapter(chapter, i)
	}

	err = e.epubExport.Write(volume.Metadata.Title + ".epub")
	if err != nil {
		return err
	}
	return nil
}

func (e EpubExporter) checkError(err error) {
	if err != nil {
		e.Callback(err.Error())
	}
}

func (e *EpubExporter) AddChapter(chapter yoinker.Chapter, i int) {
	chapterName := fmt.Sprintf("Chapter %d", i)
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
	e.epubExport.AddSection(parsedContent.String(), chapterName, "", "")
}
