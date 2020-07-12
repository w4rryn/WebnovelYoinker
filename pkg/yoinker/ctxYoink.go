package yoinker

import "github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker/book"

//CtxYoink Provides event context
type CtxYoink struct {
	Volume book.Volume
	Error  error
}
