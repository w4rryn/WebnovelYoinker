package book

// Volume represents a book that consists of chapters
type Volume struct {
	Chapters []Chapter `json:"chapters"`
	Metadata Metadata  `json:"metadata"`
}
