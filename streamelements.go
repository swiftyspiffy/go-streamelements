package streamelements

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go-streamelements/events"
	"time"
)

type SE interface {
	Connect(context.Context) error
	Disconnect() error
}

type StreamElements struct {
	config *Config

	client     *websocket.Dialer
	conn       *websocket.Conn
	pingerChan chan bool
}

func New(config *Config) SE {
	return &StreamElements{
		config: config,
		client: &websocket.Dialer{},
	}
}

func (se *StreamElements) Connect(ctx context.Context) error {
	conn, _, err := se.client.DialContext(ctx, se.config.WebsocketEndpoint, nil)
	if err != nil {
		return err
	}
	events.ClientConnected.Trigger(&events.ClientConnectedPayload{})

	se.conn = conn

	go func() {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				se.pingerChan <- true
				_ = se.conn.Close()
				fmt.Printf("ERROR: %s\n", err)
				return
			}
			events.ClientReceived.Trigger(&events.ClientReceivedPayload{
				IsText:  messageType == websocket.TextMessage,
				Message: message,
			})
			if messageType == websocket.TextMessage {
				if err = se.parseMessage(string(message)); err != nil {
					se.handleError(err)
					continue
				}
			}
		}
	}()

	return nil
}

func (se *StreamElements) Disconnect() error {
	se.pingerChan <- true
	if err := se.conn.Close(); err != nil {
		return err
	}
	events.ClientDisconnected.Trigger(&events.ClientDisconnectedPayload{})
	return nil
}

func (se *StreamElements) Send(message string) error {
	if err := se.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		return err
	}
	events.ClientSent.Trigger(&events.ClientSentPayload{
		Message: message,
	})
	return nil
}

func (se *StreamElements) startPinger(intervalMs int, timeoutMs int) {
	if err := se.Send("2"); err != nil {
		se.handleError(err)
	}

	interval := time.Duration(intervalMs) * time.Millisecond

	ticker := time.NewTicker(interval)
	se.pingerChan = make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := se.Send("2"); err != nil {
					se.handleError(err)
				}
			case <-se.pingerChan:
				return
			}
		}
	}()
}

func (se *StreamElements) handleError(err error) {
	events.ClientError.Trigger(&events.ClientErrorPayload{
		Err: err,
	})
}
