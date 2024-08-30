package events

var ClientUnknownMessage clientUnknownMessage

type ClientUnknownMessagePayload struct {
	Message string
}

type clientUnknownMessage struct {
	handlers []interface {
		UnknownMessageReceived(payload *ClientUnknownMessagePayload)
	}
}

func (e *clientUnknownMessage) Register(handler interface {
	UnknownMessageReceived(payload *ClientUnknownMessagePayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientUnknownMessage) Trigger(payload *ClientUnknownMessagePayload) {
	for _, handler := range e.handlers {
		go handler.UnknownMessageReceived(payload)
	}
}
