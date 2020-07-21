package export

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"
)

type pdfExporter struct {
}

func (p *pdfExporter) Export(metadata book.Metadata, path string, chapters []book.Chapter) string {
	const (
		fontSize        = 12
		columnWidth     = 12
		verticalPadding = 2.0
		rowHeight       = 5
		rowBufferSize   = 54
	)
	var (
		paragraphStyle = props.Text{
			Size:            fontSize,
			Align:           consts.Left,
			Top:             0,
			VerticalPadding: verticalPadding,
		}
		line       []string
		exportPath = filepath.Join(path, metadata.Title+".pdf")
	)
	m := pdf.NewMaroto(consts.Portrait, consts.A5)
	for _, chapter := range chapters {
		m.Row(20, func() {
			m.Col(columnWidth, func() {
				m.Text(chapter.ChapterName, props.Text{
					Size:            40,
					Align:           consts.Center,
					VerticalPadding: 40,
				})
			})
		})
		for _, content := range chapter.Content {
			m.Row(rowHeight, func() {
				m.Col(columnWidth, func() {
					m.Text("", paragraphStyle)
				})
			})
			switch content.(type) {
			case *book.Paragraph:
				text := content.(*book.Paragraph)
				for _, word := range strings.Fields(text.Content) {
					line = append(line, word)
					joinedLine := strings.Join(line, " ")
					if len(joinedLine) >= rowBufferSize {
						m.Row(rowHeight, func() {
							m.Col(columnWidth, func() {
								m.Text(joinedLine, paragraphStyle)
							})
						})
						line = nil
					}
				}
			}
		}
		m.AddPage()
	}
	err := m.OutputFileAndClose(exportPath)
	if err != nil {
		log.Printf("Could not save PDF: %v\n", err)
	}

	return exportPath
}

//NewPdfExporter creates a new instance of pdfExporter
func NewPdfExporter() yoinker.IExportStrategy {
	return &pdfExporter{}
}
