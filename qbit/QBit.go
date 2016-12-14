package qbit

import (
	"time"
)

type Queue  interface {
	ReceiveQueue() ReceiveQueue
	SendQueue() SendQueue
	SendQueueWithAutoFlush(interval int, duration time.Duration) SendQueue
	StartListener(listener ReceiveQueueListener) error
	Started() bool
	Stopped() bool
	Size() int
	Name() string
	Stop() error
}

type SendQueue interface {
	Send(item interface{}) (bool, error)
	FlushSends() error
	Size() int
	Name() string
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




