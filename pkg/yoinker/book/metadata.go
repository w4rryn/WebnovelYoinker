package book

//Metadata stores metadata for books
type Metadata struct {
	Website     ScrapingWebsite `json:"website" yaml:"Website"`
	Title       string          `json:"title" yaml:"Title"`
	Author      string          `json:"author" yaml:"Author"`
	Language    string          `json:"language" yaml:"Language"`
	Year        string          `json:"year" yaml:"Year"`
	Cover       string          `json:"cover" yaml:"CoverImageURL"`
	ChapterURLs []string        `json:"chapterURLs" yaml:"ChapterURLs"`
	Format      Exporters       `json:"exportFormat" yaml:"ExportFormat"`
}
