package qbit

import (
	"time"
)

type BasicQueue struct {
	channel          chan []interface{}
	batchSize        int
	channelSize      int
	pollWaitDuration time.Duration
}

func NewQueue(batchSize int, channelSize int, pollWaitDuration time.Duration) Queue {
	channel := make(chan []interface{}, channelSize)
	queue := &BasicQueue{
		channel:          channel,
		batchSize:        batchSize,
		pollWaitDuration: pollWaitDuration,
	}
	return queue
}

func (bq *BasicQueue) ReceiveQueue() ReceiveQueue {
	return NewBasicReceiveQueue(bq.pollWaitDuration, bq.channel)
}

func (bq *BasicQueue) SendQueue() SendQueue {
	return NewSendQueue(bq.channel, bq, bq.batchSize, nil)
}

func (bq *BasicQueue) SendQueueWithAutoFlush(flushDuration time.Duration) SendQueue {

	sendQueue := NewLockingSendQueue(bq.SendQueue())

	return sendQueue
}

func (bq *BasicQueue) Size() int {
	return len(bq.channel)
}
