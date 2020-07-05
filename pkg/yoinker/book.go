package yoinker

// Volume represents a book that consists of chapters
type Volume struct {
	Cover    string
	Chapters []Chapter
	Author   string
	Year     string
	Language string
	Title    string
}

// Chapter represents a chapter in a book.
// Chapter content must be valid xhtml
type Chapter struct {
	Images  []string
	Content []string
}
