package events

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"

//OnExportStartEvent gets raised when a new volume gets exportet
var OnExportStartEvent = &onExportStart{}

type onExportStart struct {
	handles []yoinker.EventHandle
}

func (e *onExportStart) Invoke(y *yoinker.CtxYoink) {
	invoke(e.handles, y)
}

func (e *onExportStart) Add(handler yoinker.EventHandle) {
	e.handles = append(e.handles, handler)
}
