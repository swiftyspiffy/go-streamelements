package events

var ClientDisconnected clientDisconnected

type ClientDisconnectedPayload struct {
}

type clientDisconnected struct {
	handlers []interface {
		Disconnected(payload *ClientDisconnectedPayload)
	}
}

func (e *clientDisconnected) Register(handler interface {
	Disconnected(payload *ClientDisconnectedPayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientDisconnected) Trigger(payload *ClientDisconnectedPayload) {
	for _, handler := range e.handlers {
		go handler.Disconnected(payload)
	}
}
