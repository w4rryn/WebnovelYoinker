package yoinker

//ExportStrategy interface to provide document export functionality
type ExportStrategy interface {
	Export(volume *Volume) error
}
