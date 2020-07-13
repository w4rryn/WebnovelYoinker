package terminal

// import (
// 	"fmt"

// 	"github.com/urfave/cli/v2"
// )

// var singleScrapeFlags = []cli.Flag{
// 	&cli.StringFlag{
// 		Name:     "url",
// 		Aliases:  []string{"u"},
// 		Required: true,
// 	},
// 	&cli.StringFlag{
// 		Name:     "website",
// 		Aliases:  []string{"w"},
// 		Required: true,
// 	},
// }

// func singleScrapeCommand(c *cli.Context) error {
// 	yoinkManager := factory.NewYoinker()
// 	volumes := yoinkManager.GetAvailableVolumes(c.String("url"), c.String("website"))

// 	n := 0
// 	for _, volume := range volumes {
// 		fmt.Println(volume.Metadata.Title)
// 		for _, chapter := range volume.Chapters {
// 			fmt.Printf("%v %v: %v\n", n, chapter.ChapterName, chapter.URL)
// 			n++
// 		}
// 	}
// 	return nil
// }
