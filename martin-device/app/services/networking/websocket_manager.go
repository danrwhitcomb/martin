package networking

import (
	"container/list"
	"encoding/json"
	"fmt"

	"github.com/revel/revel"

	"martin/martin-device/app/services/dto"
)

// Singleton variables that exist for lifecycle of app
var (
	incomingMessages = make(chan string, 1)
	manager          *WebSocketManager
)

type Connection struct {
	Id       string
	Outgoing chan<- string
}

// ClientManager is notified of changes to a manager
type ClientManager interface {
	AddObserver(chan dto.ClientEvent)
	HasClients() bool
	Send(data interface{}) error
}

// WebSocketManager (implementation of ClientManager)
// Handles messagings via a single websocket connection
type WebSocketManager struct {
	Connections      *list.List
	observers        *list.List
	incomingMessages chan map[string]interface{}
}

// GetWebSocketManager returns a singleton WebSocketManager
func GetWebSocketManager() *WebSocketManager {
	if manager == nil {
		manager = &WebSocketManager{Connections: list.New(),
			observers:        list.New(),
			incomingMessages: make(chan map[string]interface{}, 1)}

		go manager.run()
	}

	return manager
}

// AddObserver adds an observer
func (c *WebSocketManager) AddObserver(channel chan dto.ClientEvent) {
	c.observers.PushBack(channel)
}

// Check if clients are connected or not
func (c *WebSocketManager) HasClients() bool {
	return c.Connections.Len() != 0
}

// Send data to the clients
func (c *WebSocketManager) Send(data interface{}) error {

	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Send message to all clients
	for conn := c.Connections.Front(); conn != nil; conn = conn.Next() {
		conn.Value.(Connection).Outgoing <- string(jsonString)
	}

	return nil
}

func (c *WebSocketManager) RegisterConnection(id string, outChannel chan<- string) chan<- map[string]interface{} {
	c.Connections.PushBack(Connection{Id: id, Outgoing: outChannel})
	defer c.notifyObservers(dto.Connected, map[string]interface{}{"id": id})

	return c.incomingMessages
}

func (c *WebSocketManager) RemoveConnection(id string) {

	for conn := c.Connections.Front(); conn != nil; conn = conn.Next() {
		if conn.Value.(Connection).Id == id {
			c.Connections.Remove(conn)
		}
	}

	c.notifyObservers(dto.Disconnected, map[string]interface{}{"id": id})
}

// Listens to incoming messaging channel
// then reads from received channel and notifies observers
func (c *WebSocketManager) run() {
	for {
		message := <-c.incomingMessages
		fmt.Println(message)
		c.notifyObservers(dto.SentMessage, message)
	}
}

// Helper function for notifying observers of some event
func (c *WebSocketManager) notifyObservers(event dto.EventType, message map[string]interface{}) {
	revel.AppLog.Infof("Notifying Observers: %d", event)
	for obs := c.observers.Front(); obs != nil; obs = obs.Next() {
		obs.Value.(chan dto.ClientEvent) <- dto.ClientEvent{EventType: event, Data: message}
	}
}

func init() {
	GetWebSocketManager()
}
