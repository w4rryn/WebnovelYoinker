package yoinker

//BookMetadata stores metadata for books
type BookMetadata struct {
	WebsiteURL  string   `json:"website" yaml:"Website"`
	Title       string   `json:"title" yaml:"Title"`
	Author      string   `json:"author" yaml:"Author"`
	Language    string   `json:"language" yaml:"Language"`
	Year        string   `json:"year" yaml:"Year"`
	Cover       string   `json:"cover" yaml:"CoverImageURL"`
	ChapterURLs []string `json:"chapterURLs" yaml:"ChapterURLs"`
	Format      string   `json:"exportFormat" yaml:"ExportFormat"`
}
