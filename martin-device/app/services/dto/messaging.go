package dto

import (
	"fmt"
)

/**
 * Event system for managing connection lifecycle
 */
type EventType int

func (e EventType) String() string {
	if e == Connected {
		return "Connected"
	} else if e == Disconnected {
		return "Disconnected"
	} else if e == SentMessage {
		return "SentMessage"
	}

	return e.String()
}

const (
	Connected EventType = iota + 1
	Disconnected
	SentMessage
)

type ClientEvent struct {
	EventType EventType
	Data      map[string]interface{}
}

// String converter for Event
func (c *ClientEvent) String() string {
	return fmt.Sprintf("Event{EventType: %s, Data: %s}", c.EventType.String(), c.Data)
}

//

type ClientNotification struct {
}
