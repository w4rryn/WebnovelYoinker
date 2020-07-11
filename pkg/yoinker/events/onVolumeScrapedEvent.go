package events

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"

//OnVolumeScrapedEvent raised when a volume was scraped successfully
var OnVolumeScrapedEvent = &onVolumeScraped{}

type onVolumeScraped struct {
	handles []yoinker.EventHandle
}

func (c *onVolumeScraped) Invoke(y *yoinker.CtxYoink) {
	invoke(c.handles, y)
}

func (c *onVolumeScraped) Add(handler yoinker.EventHandle) {
	c.handles = append(c.handles, handler)
}
