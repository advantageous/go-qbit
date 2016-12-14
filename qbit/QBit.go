package qbit

import (
	"time"
)

type Queue interface {
	ReceiveQueue() ReceiveQueue
	SendQueue() SendQueue
	SendQueueWithAutoFlush(duration time.Duration) SendQueue
	Stopped() bool
	Size() int
	Stop() error
}

type SendQueue interface {
	Send(item interface{}) error
	FlushSends() error
	Size() int
}

type ReceiveQueue interface {
	PollWait() interface{}
	Poll() interface{}
	Take() interface{}
	TakeBatch() []interface{}
	ReadBatch() []interface{}
	ReadBatchWait() []interface{}
	Size() int
}

type ReceiveQueueListener interface {
	Init()
	Receive(item interface{})
	Empty()
	Limit()
	Shutdown()
	Idle()
	StartBatch()
}
