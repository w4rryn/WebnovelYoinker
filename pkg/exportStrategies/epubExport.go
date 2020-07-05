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
	Callback func(s string)
}

//Export exports a valume as epub
func (e *EpubExporter) Export(volume *yoinker.Volume) error {
	epubExport := epub.NewEpub(volume.Title)
	coverImage, err := epubExport.AddImage(volume.Cover, "")
	if err != nil {
		return err
	}
	epubExport.SetCover(coverImage, "")
	epubExport.SetAuthor(volume.Author)
	epubExport.SetLang(volume.Language)

	for i, chapter := range volume.Chapters {
		chapterName := fmt.Sprintf("Chapter %d", i+1)
		var parsedContent strings.Builder
		for _, paragraph := range chapter.Content {
			switch paragraph.(type) {
			case *yoinker.PageImage:
				pageImage := paragraph.(*yoinker.PageImage)
				imagePath, err := epubExport.AddImage(pageImage.Image, "")
				e.checkError(err)
				content := fmt.Sprintf("<p style=\"page-break-before: always\"><img src=\"%v\" width=\"%v\" height=\"%v\"/></p>", imagePath, pageImage.Width, pageImage.Height)
				parsedContent.WriteString(content)
			case *yoinker.Paragraph:
				par := paragraph.(*yoinker.Paragraph)
				var content string
				if par.Content == "" {
					content = "<p> </p>"
				} else {
					content = fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
				}
				parsedContent.WriteString(content)
			}
		}
		epubExport.AddSection(parsedContent.String(), chapterName, "", "")
	}
	err = epubExport.Write(volume.Title + ".epub")
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
