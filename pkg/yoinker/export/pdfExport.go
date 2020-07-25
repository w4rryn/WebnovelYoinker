package export

import (
	"bytes"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//PdfExporter text  as PDF
type PdfExporter struct {
}

//Export exports given strings as PDF
func (p *PdfExporter) Export(metadata book.Metadata, path string, chapters []book.Chapter) string {
	var (
		body       = createBookHTML(chapters, metadata.Title)
		outputPath = filepath.Join(path, metadata.Title+".pdf")
		pdfg       = wkhtmltopdf.NewPDFPreparer()
	)
	file, err := os.Create(metadata.Title + "_dump.html")
	p.checkError(err)

	file.WriteString(body)
	defer file.Close()

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	pdfg.TOC.Include = true
	page := wkhtmltopdf.NewPageReader(strings.NewReader(body))
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(0)
	page.Zoom.Set(0.95)

	pdfg.AddPage(page)

	jsonBytes, err := pdfg.ToJSON()
	p.checkError(err)

	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	p.checkError(err)

	err = pdfgFromJSON.Create()
	p.checkError(err)

	err = pdfgFromJSON.WriteFile(outputPath)
	p.checkError(err)

	return outputPath
}

//createBookHTML converts a chapter array to a html page
func createBookHTML(volume []book.Chapter, title string) string {
	var body strings.Builder
	for _, chapter := range volume {
		isText := false
		body.WriteString(fmt.Sprintf("<h1>%v</h1>", html.EscapeString(chapter.ChapterName)))
		for _, par := range chapter.Content {
			switch par.(type) {
			case *book.PageImage:
				if isText {
					body.WriteString("</div>")
					isText = false
				}
				pageImage := par.(*book.PageImage)
				content := fmt.Sprintf("<div class=\"width pc\"><img class=\"calibre1\" alt=\"image\" src=\"%v\" width=\"%v\" height=\"%v\"/></div>",
					pageImage.Image, pageImage.Width, pageImage.Height)
				body.WriteString(content)

			case *book.Paragraph:
				if !isText {
					body.WriteString("<div class=\"text-body\"")
					isText = true
				}
				par := par.(*book.Paragraph)
				content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
				body.WriteString(content)
			}
		}
		if isText {
			body.WriteString("</div>")
		}

	}
	return fmt.Sprintf("<!doctype html>"+
		"<meta charset=\"utf-8\"/>"+
		"<html>"+
		"<head>"+
		"<style>"+
		".text-body{font-size: 16px;}"+
		"</style>"+
		"<title><h1>%v</h1></title>"+
		"</head>"+
		"<body>"+
		"%v"+
		"</body>"+
		"</html>", title, body.String())
}

func (p PdfExporter) checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

//NewPdfExporter creates a new instance of pdfExporter
func NewPdfExporter() *PdfExporter {
	return &PdfExporter{}
}
