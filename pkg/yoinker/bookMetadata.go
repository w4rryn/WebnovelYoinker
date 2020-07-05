package yoinker

//BookMetadata stores metadata for books
type BookMetadata struct {
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Year     string `json:"year"`
	Language string `json:"language"`
	Title    string `json:"title"`
}
