package events

var ClientError clientError

type ClientErrorPayload struct {
	Err error
}

type clientError struct {
	handlers []interface {
		Errored(payload *ClientErrorPayload)
	}
}

func (e *clientError) Register(handler interface {
	Errored(payload *ClientErrorPayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientError) Trigger(payload *ClientErrorPayload) {
	for _, handler := range e.handlers {
		go handler.Errored(payload)
	}
}
