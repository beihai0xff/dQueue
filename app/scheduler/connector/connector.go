package connector

import (
	"context"

	"github.com/beihai0xff/pudding/types"
)

// nolint:lll
//go:generate mockgen -destination=../../../test/mock/connector_mock.go -package=mock github.com/beihai0xff/pudding/app/scheduler/connector RealTimeConnector

// RealTimeConnector is a connector which can send messages to the realtime queue
// the realtime queue can store or consume messages in realtime
type RealTimeConnector interface {
	// Produce produce a Message to the queue in real time
	Produce(ctx context.Context, msg *types.Message) error
	// NewConsumer new a consumer to consume Messages from the realtime queue in background
	NewConsumer(topic, group string, batchSize int, fn types.HandleMessage) error
	// Close the queue
	Close() error
}
