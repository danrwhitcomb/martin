package controllers

import (
	"encoding/json"
	"strconv"
	"sync/atomic"

	"martin/martin-device/app/services/networking"

	"github.com/revel/revel"
)

// **********  CONTROLLER *****************
// Handles accepting websocket connections, then sending and
// receiving messages
type WebSocketController struct {
	*revel.Controller
}

var count uint64 = 0
var log = revel.AppLog

// Connection handles incoming websocket connections
func (c *WebSocketController) Connection(ws revel.ServerWebSocket) revel.Result {

	manager := networking.GetWebSocketManager()

	// Register the new connection with the manager
	// and setup incoming and outgoing channels
	id := strconv.FormatUint(atomic.AddUint64(&count, 1), 10)
	outgoingChannel := make(chan string, 1)
	incomingChannel := manager.RegisterConnection(id, outgoingChannel)
	defer manager.RemoveConnection(id)
	defer close(outgoingChannel)

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan map[string]interface{})
	go func() {
		var msg map[string]interface{}
		for {
			err := ws.MessageReceiveJSON(&msg)
			switch err.(type) {
			case *json.SyntaxError:
				log.Debugf("JSON syntax error: %s", err)
				continue
			case error:
				log.Debug("Error occurred, closing connection")
				close(newMessages)
				return
			}

			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case outgoingMessage := <-outgoingChannel:
			if ws.MessageSendJSON(&outgoingMessage) != nil {
				// They disconnected
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, notify the oberservers
			incomingChannel <- msg
		}
	}
}

func NotifyNotAllowed(ws revel.ServerWebSocket) {

}
