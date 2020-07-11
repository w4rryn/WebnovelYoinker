package events

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"

//OnExportFinishedEvent Raised, when export of book is finished
var OnExportFinishedEvent = &onExportFinished{}

//onExportFinished Raised, when export of book is finished
type onExportFinished struct {
	handler []yoinker.EventHandle
}

//Invoke invokes the event
func (e *onExportFinished) Invoke(c *yoinker.CtxYoink) {
	invoke(e.handler, c)
}

//Add registers a new event handler to this event
func (e *onExportFinished) Add(handle yoinker.EventHandle) {
	e.handler = append(e.handler, handle)
}
