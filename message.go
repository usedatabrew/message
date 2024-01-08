package message

import "sync"

// Message represents a Blink message and contains raw data object
// linked to the data stream and event.
type Message struct {
	lock   sync.RWMutex
	Event  Event  `json:"event" yaml:"event"`
	Stream string `json:"stream" yaml:"stream"`
	Data   Data   `json:"data" yaml:"data"`
}

func NewMessage(event Event, stream string, data []byte) *Message {
	return &Message{
		Event:  event,
		Stream: stream,
		Data:   NewData(data),
	}
}

func (m *Message) GetStream() string {
	return m.Stream
}

func (m *Message) GetEvent() Event {
	return m.Event
}

func (m *Message) AsJSONString() string {
	return m.Data.packet.JSON()
}
