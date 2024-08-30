package events

var ClientAuthResponse clientAuthResponse

type ClientAuthResponsePayload struct {
	Successful bool
}

type clientAuthResponse struct {
	handlers []interface {
		AuthResponseReceived(payload *ClientAuthResponsePayload)
	}
}

func (e *clientAuthResponse) Register(handler interface {
	AuthResponseReceived(payload *ClientAuthResponsePayload)
}) {
	e.handlers = append(e.handlers, handler)
}

func (e *clientAuthResponse) Trigger(payload *ClientAuthResponsePayload) {
	for _, handler := range e.handlers {
		go handler.AuthResponseReceived(payload)
	}
}
