package events

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"

func invoke(handler []yoinker.EventHandle, c *yoinker.CtxYoink) {
	for _, handle := range handler {
		if handle != nil {
			go handle(c)
		}
	}
}
