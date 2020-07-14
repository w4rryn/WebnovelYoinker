package export

import (
	"errors"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

//GetExporter gets instance of a specific export strategy
func GetExporter(export book.Exporters) (yoinker.IExportStrategy, error) {
	switch export {
	case book.EPUB:
		return NewEpubExporter(), nil
	}
	return nil, errors.New("Exporter not supported")
}
