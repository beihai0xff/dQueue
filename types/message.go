package types

import (
	"context"
	"encoding/json"
)

type HandleMessage func(ctx context.Context, msg *Message) error

// Message 消息
type Message struct {
	Topic        string // Message Topic
	Key          string // Message Key
	Payload      []byte // Message Payload
	DeliverAfter int64  // Message DeliverAfter time (Seconds)
	DeliverAt    int64  // Message Deliver timestamp（now + DeliverAfter, Unix Timestamp, Seconds）
}

func GetMessageFromJSON(j []byte) (*Message, error) {
	var m Message
	if err := json.Unmarshal(j, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
