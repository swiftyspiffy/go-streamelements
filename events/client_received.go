package events

var ClientReceived clientReceived

type ClientReceivedPayload struct {
	IsText  bool
	Message []byte
}

type clientReceived struct {
	handlers []interface {
		Received(payload *ClientReceivedPayload)
	}
}

func (e *clientReceived) Register(handler interface {
	Received(payload *ClientReceivedPayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientReceived) Trigger(payload *ClientReceivedPayload) {
	for _, handler := range e.handlers {
		go handler.Received(payload)
	}
}
