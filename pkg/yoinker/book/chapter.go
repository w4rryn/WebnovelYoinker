package book

// Chapter represents a Chapter in a book.
type Chapter struct {
	ChapterName string       `json:"chapterName"`
	Images      []string     `json:"images"`
	Content     []PageObject `json:"content"`
	URL         string       `json:"url"`
}
