package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//IExportStrategy interface to provide document export functionality
type IExportStrategy interface {
	Export(metadata book.Metadata, path string, chapters []book.Chapter) string
}
