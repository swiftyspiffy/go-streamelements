package events

var ClientConnected clientConnected

type ClientConnectedPayload struct {
}

type clientConnected struct {
	handlers []interface {
		Connected(payload *ClientConnectedPayload)
	}
}

func (e *clientConnected) Register(handler interface {
	Connected(payload *ClientConnectedPayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientConnected) Trigger(payload *ClientConnectedPayload) {
	for _, handler := range e.handlers {
		go handler.Connected(payload)
	}
}
