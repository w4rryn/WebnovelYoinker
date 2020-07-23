package factories

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/export"
)

//GetExporter gets instance of a specific export strategy
func GetExporter(exporter book.Exporters) (yoinker.IExportStrategy, error) {
	switch exporter {
	case book.EPUB:
		return export.NewEpubExporter(), nil
	case book.PDF:
		return export.NewPdfExporter(), nil
	}
	return nil, errors.New("Exporter not supported")
}
