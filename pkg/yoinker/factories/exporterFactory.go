package factories

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/export"
)

//GetExporter gets instance of a specific export strategy
func GetExporter(exporter string) (yoinker.IExportStrategy, error) {
	switch exporter {
	case string(export.EPUB):
		return export.NewEpubExporter(), nil
	case string(export.PDF):
		return export.NewPdfExporter(), nil
	}
	return nil, errors.New("Exporter not supported")
}
