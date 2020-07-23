package export

import (
	"bytes"
	"fmt"
	"html"
	"log"
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
	var body strings.Builder
	for _, chapter := range chapters {
		body.WriteString(fmt.Sprintf("<h1>%v</h1>", html.EscapeString(chapter.ChapterName)))
		for _, par := range chapter.Content {
			switch par.(type) {
			case *book.PageImage:
				pageImage := par.(*book.PageImage)
				content := fmt.Sprintf("<div>"+
					"<img src=\"%v\" width=\"%v\" height=\"%v\""+
					"</div>",
					pageImage.Image, pageImage.Width, pageImage.Height)
				body.WriteString(content)

			case *book.Paragraph:
				par := par.(*book.Paragraph)
				content := fmt.Sprintf("<p>%v</p>", html.EscapeString(par.Content))
				body.WriteString(content)
			}
		}
	}
	var html = fmt.Sprintf("<!doctype html><meta charset=\"utf-8\"/><html><head><title><h1>%v</h1></title></head>%v</body></html>", metadata.Title, body.String())

	var outputPath = filepath.Join(path, metadata.Title+".pdf")
	// Client code
	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	pdfg.AddPage(page)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.WriteFile(outputPath)
	if err != nil {
		log.Println(err)
	}

	return outputPath
}

//NewPdfExporter creates a new instance of pdfExporter
func NewPdfExporter() *PdfExporter {
	return &PdfExporter{}
}
