package qbit

import (
	"time"
)

type Queue interface {
	ReceiveQueue() ReceiveQueue
	SendQueue() SendQueue
	Size() int
}

type QueueManager interface {
	SendQueueWithAutoFlush(duration time.Duration) SendQueue
	Stopped() bool
	Stop() error
	Queue() Queue
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
	TakeBatch()  *ChannelBuffer
	ReadBatch() *ChannelBuffer
	ReadBatchWait()  *ChannelBuffer
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

type ChannelBuffer struct {
	Buffer []interface{}
	Index  int
}
