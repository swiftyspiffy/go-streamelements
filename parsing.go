package streamelements

import (
	"encoding/json"
	"fmt"
	"go-streamelements/events"
	"strings"
)

func (se *StreamElements) parseMessage(wsMsg string) error {
	// identify message type by prefixed number
	fmt.Printf("%s\n", wsMsg)
	var msg string
	if strings.Contains(wsMsg, "\"") {
		number := strings.Split(wsMsg, "\"")[0 : len(strings.Split(wsMsg, "\"")[0])-1]
		msg = wsMsg[len(number):]
	}

	// send auth if needed
	if strings.HasPrefix(wsMsg, "40") {
		return se.Send(fmt.Sprintf("42[\"authenticate\",{{\"method\":\"jwt\",\"token\":\"%s\"}}]", se.config.JWT))
	}

	// send ping if needed
	if strings.HasPrefix(wsMsg, "0{\"sid\"") {
		metadata, err := parseSessionMetadata(msg)
		if err != nil {
			return err
		}
		// start pinger
		se.startPinger(metadata.PingInterval, metadata.PingTimeout)
		return nil
	}

	// handle successful auth
	if strings.HasPrefix(wsMsg, "42[\"authenticate\"") {
		events.ClientAuthResponse.Trigger(&events.ClientAuthResponsePayload{
			Successful: true,
		})
		return nil
	}

	// handle unsuccessful auth
	if strings.HasPrefix(wsMsg, "42[\"unauthorized\"") {
		events.ClientAuthResponse.Trigger(&events.ClientAuthResponsePayload{
			Successful: false,
		})
		return nil
	}

	// handle complex object
	if strings.HasPrefix(wsMsg, "42[\"event\",{\"type\"") {
		return handleComplexObject(msg)
	}

	// handle simple updates
	if strings.HasPrefix(wsMsg, "42[\"event:update\",{\"name\"") {
		return handleSimpleUpdates(msg)
	}

	// unknown message, notify consumer if they're tracking
	events.ClientUnknownMessage.Trigger(&events.ClientUnknownMessagePayload{
		Message: msg,
	})
	return nil
}

type sessionMetadata struct {
	SessionId    string `json:"sid,omitempty"`
	PingInterval int    `json:"pingInterval,omitempty"`
	PingTimeout  int    `json:"pingTimeout,omitempty"`
	MaxPayload   int    `json:"maxPayload,omitempty"`
}

func parseSessionMetadata(msg string) (*sessionMetadata, error) {
	var metadata sessionMetadata
	if err := json.Unmarshal([]byte(msg), &metadata); err != nil {
		return nil, err
	}
	return &metadata, nil
}

type eventPayload struct {
	Event    string `json:"event"`
	Provider string `json:"provider,omitempty"`
	Data     string `json:"data"`
}

func handleComplexObject(msg string) error {
	var event eventPayload
	if err := json.Unmarshal([]byte(msg), &event); err != nil {
		return err
	}

	fmt.Printf("event: %s, data: %s\n", event.Event, event.Data)

	switch event.Event {
	case "follow":

	case "cheer":

	case "host":

	case "tip":

	case "subscriber":

	default:
		events.ClientUnknownMessage.Trigger(&events.ClientUnknownMessagePayload{
			Message: msg,
		})
	}
	return nil
}

func handleSimpleUpdates(msg string) error {
	fmt.Printf("msg: %s\n", msg)
	return nil
}

func parseFollow() {

}

func parseCheer() {

}

func parseHost() {

}

func parseTip() {

}

func parseSubscriber() {

}
