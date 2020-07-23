package export

import (
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bmaupin/go-epub"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

var extensionMediaTypes = map[string]string{
	".css":   "text/css",
	".gif":   "image/gif",
	".jpeg":  "image/jpeg",
	".jpg":   "image/jpeg",
	".otf":   "application/vnd.ms-opentype",
	".png":   "image/png",
	".svg":   "image/svg+xml",
	".ttf":   "application/font-sfnt",
	".woff":  "application/font-woff",
	".woff2": "font/woff2",
}

//EpubExporter exports a volume a epub
type EpubExporter struct {
	epubExport *epub.Epub
	fileDump   *os.File
}

//Export exports a valume as epub
func (e *EpubExporter) Export(metadata book.Metadata, path string, chapters []book.Chapter) string {
	var (
		coverImage string
		waiter     sync.WaitGroup
		exportPath = filepath.Join(path, metadata.Title+".epub")
		err        error
		cssPath    string
	)
	// e.fileDump, err = os.Create(filepath.Join(path, fmt.Sprintf("volume_dump_%v_.txt", metadata.Title)))
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer e.fileDump.Close()

	e.epubExport = epub.NewEpub(metadata.Title)
	cssPath, err = e.epubExport.AddCSS("https://raw.githubusercontent.com/lethal-bacon0/WebnovelYoinker/master/assets/ebookstyle.css", "stylesheet.css")
	if err != nil {
		log.Println(err)
	}
	waiter.Add(1)
	go func() {
		coverImage, err = e.epubExport.AddImage(metadata.Cover, "")
		if err != nil {
			coverImage = ""
		} else {
			e.epubExport.SetCover(coverImage, "")
		}
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
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic while exporting %v: %v\n", metadata.Title, r)
		}
	}()
	err = e.epubExport.Write(exportPath)
	if err != nil {
		log.Println(err)
	}

	return exportPath
}

func (e *EpubExporter) addChapter(chapter book.Chapter) string {
	var (
		parsedContent strings.Builder
	)

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
			// e.fileDump.WriteString(fmt.Sprintf("%v\n", par.Content))
			content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
			parsedContent.WriteString(content)
		}
	}
	return parsedContent.String()
}

//NewEpubExporter creates a new epub exporter
func NewEpubExporter() *EpubExporter {
	return &EpubExporter{}
}
