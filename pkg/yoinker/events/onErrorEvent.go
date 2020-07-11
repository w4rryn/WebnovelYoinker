package events

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"

//OnErrorEvent raised on error
var OnErrorEvent = &onError{}

type onError struct {
	handles []yoinker.EventHandle
}

func (e *onError) Invoke(y *yoinker.CtxYoink) {
	invoke(e.handles, y)
}

func (e *onError) Add(handler yoinker.EventHandle) {
	e.handles = append(e.handles, handler)
}
