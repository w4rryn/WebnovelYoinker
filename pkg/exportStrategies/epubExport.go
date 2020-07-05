package exportstrategies

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/bmaupin/go-epub"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

//EpubExporter exports a volume a epub
type EpubExporter struct {
}

//Export exports a valume as epub
func (e *EpubExporter) Export(volume *yoinker.Volume) error {
	epubExport := epub.NewEpub(volume.Title)
	coverImage, err := epubExport.AddImage(volume.Cover, "")
	if err != nil {
		return err
	}
	epubExport.SetCover(coverImage, "")
	epubExport.SetAuthor(volume.Author)
	epubExport.SetLang(volume.Language)

	for i, chapter := range volume.Chapters {
		var localImages []string
		for _, imageURL := range chapter.Images {
			epubImage, err := epubExport.AddImage(imageURL, "")
			if err != nil {
				return err
			}
			localImages = append(localImages, epubImage)
		}
		chapter.Images = localImages
		chapterName := fmt.Sprintf("Chapter %d", i+1)
		chapterTemplate := template.New("Template_1")
		chapterTemplate, _ = chapterTemplate.Parse(chapter.Content)
		var stringBuffer bytes.Buffer
		err := chapterTemplate.Execute(&stringBuffer, chapter)
		if err != nil {
			return err
		}
		fmt.Println(stringBuffer.String())
		epubExport.AddSection(stringBuffer.String(), chapterName, "", "")
	}
	err = epubExport.Write(volume.Title)
	if err != nil {
		return err
	}
	return nil
}
