package yc

//ExportStrategy interface to provide document export functionality
type ExportStrategy interface {
	Export(metadata BookMetadata, path string, chapter <-chan Chapter) string
}
