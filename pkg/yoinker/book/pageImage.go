package book

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
