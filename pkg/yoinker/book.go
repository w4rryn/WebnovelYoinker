package yoinker

// volume represents a book that consists of chapters
type volume struct {
	Chapters []chapter    `json:"chapters"`
	Metadata BookMetadata `json:"metadata"`
}

// Chapter represents a chapter in a book.
type chapter struct {
	ChapterName string       `json:"chapterName"`
	Images      []string     `json:"images"`
	Content     []pageObject `json:"content"`
}

var pageImageID uint

//PageImage is used to store an image on a page
type pageImage struct {
	ID     uint   `json:"id"`
	Image  string `json:"image"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

//GetID gets the ID of the current instance
func (p *pageImage) GetID() uint {
	if p.ID <= 0 {
		pageImageID++
		p.ID = pageImageID
		return pageImageID
	}
	return p.ID
}

//newPageImage creates a new page
func newPageImage(height int, width int, imagePath string) *pageImage {
	return &pageImage{
		Height: height,
		Image:  imagePath,
		Width:  width,
	}
}

var paragraphID uint

//Paragraph represents a paragraph in a chapter
type paragraph struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

//GetID gets the id of the current paragraph
func (p *paragraph) GetID() uint {
	if p.ID <= 0 {
		paragraphID++
		p.ID = paragraphID
		return paragraphID
	}
	return p.ID
}

//NewParagraph creates a new paragraph
func newParagraph(content string) *paragraph {
	return &paragraph{
		Content: content,
	}
}

//PageObject is an arbitary object on a page
type pageObject interface {
	GetID() uint
}
