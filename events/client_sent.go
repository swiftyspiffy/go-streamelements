package events

var ClientSent clientSent

type ClientSentPayload struct {
	Message string
}

type clientSent struct {
	handlers []interface {
		MessageSent(payload *ClientSentPayload)
	}
}

func (e *clientSent) Register(handler interface {
	MessageSent(payload *ClientSentPayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientSent) Trigger(payload *ClientSentPayload) {
	for _, handler := range e.handlers {
		go handler.MessageSent(payload)
	}
}
