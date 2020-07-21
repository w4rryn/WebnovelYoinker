package book

var paragraphID uint

//Paragraph represents a Paragraph in a chapter
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
