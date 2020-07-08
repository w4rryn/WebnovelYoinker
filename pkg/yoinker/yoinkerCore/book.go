package yc

// Volume represents a book that consists of chapters
type Volume struct {
	Chapters []Chapter    `json:"chapters"`
	Metadata BookMetadata `json:"metadata"`
}

// Chapter represents a chapter in a book.
type Chapter struct {
	ChapterName string       `json:"chapterName"`
	Images      []string     `json:"images"`
	Content     []PageObject `json:"content"`
}

var pageImageID uint

//PageImage is used to store an image on a page
type PageImage struct {
	ID     uint   `json:"id"`
	Image  string `json:"image"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

//GetID gets the ID of the current instance
func (p *PageImage) GetID() uint {
	if p.ID <= 0 {
		pageImageID++
		p.ID = pageImageID
		return pageImageID
	}
	return p.ID
}

//NewPageImage creates a new page
func NewPageImage(height int, width int, imagePath string) *PageImage {
	return &PageImage{
		Height: height,
		Image:  imagePath,
		Width:  width,
	}
}

var paragraphID uint

//Paragraph represents a paragraph in a chapter
type Paragraph struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

//GetID gets the id of the current paragraph
func (p *Paragraph) GetID() uint {
	if p.ID <= 0 {
		paragraphID++
		p.ID = paragraphID
		return paragraphID
	}
	return p.ID
}

//NewParagraph creates a new paragraph
func NewParagraph(content string) *Paragraph {
	return &Paragraph{
		Content: content,
	}
}

//PageObject is an arbitary object on a page
type PageObject interface {
	GetID() uint
}
