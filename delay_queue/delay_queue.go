package delay_queue

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/beihai0xff/pudding/delay_queue/broker/redis_broker"
	"github.com/beihai0xff/pudding/pkg/configs"
	"github.com/beihai0xff/pudding/types"
)

type DelayQueue interface {
	// Produce produce a Message to DelayQueue
	Produce(ctx context.Context, partition string, msg *types.Message) error
	// Consume New a consumer to consume Messages from the queue
	Consume(ctx context.Context, partition string, batchSize int64, fn types.HandleMessage) error
	// Close the queue
	Close() error
}

type RealTimeQueue interface {
	// Produce produce a Message to the queue in real time
	Produce(ctx context.Context, msg *types.Message) error
	// NewConsumer consume Messages from the queue in real time
	NewConsumer(topic, group, consumerName string, batchSize int, fn func(msg *types.Message) error)
	// Close the queue
	Close() error
}

type Queue struct {
	delay    DelayQueue
	realtime RealTimeQueue
	c        *configs.DelayQueueConfig
}

func NewQueue() *Queue {
	return &Queue{
		delay:    NewDelayQueue(),
		realtime: NewRealTimeQueue(),
		c:        configs.GetDelayQueueConfig(),
	}
}

func (q *Queue) Produce(ctx context.Context, msg *types.Message) error {
	// if ReadyTime is set, use it
	// otherwise use current time
	if msg.ReadyTime <= 0 {
		if msg.Delay == 0 {
			return errors.New("message delay must be greater than 0")
		}
		msg.ReadyTime = time.Now().Unix() + msg.Delay
	} else {
		if time.Unix(msg.ReadyTime, 0).Before(time.Now()) {
			return errors.New("ReadyTime must be greater than the current time")
		}
	}

	if msg.Key == "" {
		msg.Key = uuid.NewString()
	}

	return q.delay.Produce(ctx, q.getPartition(msg.ReadyTime), msg)
}

func (q *Queue) getPartition(readyTime int64) string {
	startAt := (readyTime / q.c.PartitionInterval) * q.c.PartitionInterval
	endAt := startAt + q.c.PartitionInterval
	return fmt.Sprintf("%d-%d", startAt, endAt)
}

func (q *Queue) Consume(quit chan int) error {
	for {
		select {
		case <-quit:
			break
		default:
			partition := q.getPartition(time.Now().Unix())
			q.delay.Consume(context.Background(), partition, 10, q.moveMsgToRealTimeQueue)

		}
	}

}

func (q *Queue) moveMsgToRealTimeQueue(msg *types.Message) error {
	return q.realtime.Produce(context.Background(), msg)
}

// NewDelayQueue create a new DelayQueue
func NewDelayQueue() DelayQueue {
	return redis_broker.NewDelayQueue()
}

// NewRealTimeQueue create a new RealTimeQueue
func NewRealTimeQueue() RealTimeQueue {
	return nil
}
