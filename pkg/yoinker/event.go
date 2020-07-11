package yoinker

//EventHandle event handler method, get raised on event
type EventHandle func(y *CtxYoink)

//Event provides functionality for event driven programming
type Event interface {
	Invoke(y *CtxYoink)
	Add(handler EventHandle)
}
