package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//YoinkManager Provides Functionality to yoink Webnovels and Webtoons
type YoinkManager interface {
	StartYoink(metadata book.Metadata, exportPath string)
	GetAvailableVolumes(url string, website string) []book.Volume
}
