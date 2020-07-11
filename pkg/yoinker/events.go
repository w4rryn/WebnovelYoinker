package yoinker

var (
	OnExportFinished []yoinkerScrapeEvent
	OnExportStart    []yoinkerScrapeEvent
	OnScrapeStart    []yoinkerScrapeEvent
	OnChapterScraped []yoinkerScrapeEvent
)

type yoinkerScrapeEvent func(s string)

func invokeYoinkerScrapeEvent(events []yoinkerScrapeEvent, s string) {
	go func() {
		for _, event := range events {
			if event != nil {
				event(s)
			}
		}
	}()
}

//OnError holds events that sould be raised when an error is thrown
var OnError []func(err error)

func invokeError(err error) {
	go func() {
		for _, _error := range OnError {
			if _error != nil {
				_error(err)
			}
		}
	}()
}
