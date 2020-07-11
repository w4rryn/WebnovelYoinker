package yoinker

//YoinkManager Provides Functionality to yoink Webnovels and Webtoons
type YoinkManager interface {
	StartYoink(metadata BookMetadata, exportPath string)
}
