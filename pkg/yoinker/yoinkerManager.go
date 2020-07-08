package yoinker

import yc "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/yoinkerCore"

//YoinkManager Provides Functionality to yoink Webnovels and Webtoons
type YoinkManager interface {
	StartYoink(metadata yc.BookMetadata)
}
