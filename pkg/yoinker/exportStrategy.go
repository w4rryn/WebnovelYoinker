package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//ExportStrategy interface to provide document export functionality
type ExportStrategy interface {
	Export(metadata book.Metadata, path string, chapter <-chan book.Chapter) string
}
