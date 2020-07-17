package export

import (
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

type pdfExporter struct {
}

func (p *pdfExporter) Export(metadata book.Metadata, path string, chapters []book.Chapter) string {
	return "nil"
}

//NewPdfExporter creates a new instance of pdfExporter
func NewPdfExporter() yoinker.IExportStrategy {
	return &pdfExporter{}
}
