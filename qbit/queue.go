package qbit

import (
	"sync/atomic"
	"time"
	"errors"
)

type  BasicQueue struct {
	channel          chan []interface{}
	batchSize        int
	channelSize      int
	limit            int
	name             string
	pollWaitDuration time.Duration
	started          int64
}

func NewQueue(batchSize int, channelSize int, limit int, name string, pollWaitDuration time.Duration) Queue {
	channel := make(chan []interface{}, channelSize)
	return &BasicQueue{
		channel: channel,
		batchSize: batchSize,
		limit: limit,
		name: name,
		pollWaitDuration: pollWaitDuration,
	}
}

func (bq *BasicQueue) ReceiveQueue() ReceiveQueue {
	return NewBasicReceiveQueue(bq.pollWaitDuration, bq.channel)
}

func (bq *BasicQueue) SendQueue() SendQueue {
	return NewSendQueue(bq.channel, bq, bq.batchSize, nil)
}

func (bq *BasicQueue) SendQueueWithAutoFlush(batchSize int, flushDuration time.Duration) SendQueue {
	return nil
}

func (bq *BasicQueue)  StartListener(listener ReceiveQueueListener) error {
	var err error

	if bq.Started() {
		err = errors.New("Queue already started")
	} else if atomic.CompareAndSwapInt64(&bq.started, 0, 1) {
		go manageQueue(bq.limit, bq, bq.ReceiveQueue(), listener)
	}
	return err
}

func (bq *BasicQueue)  Started() bool {
	started := atomic.LoadInt64(&bq.started)
	return started == 1
}

func (bq *BasicQueue)  Stopped() bool {
	started := atomic.LoadInt64(&bq.started)
	return started == 0
}

func (bq *BasicQueue)  Size() int {
	return len(bq.channel)
}

func (bq *BasicQueue)  Name() string {
	return bq.name
}

func (bq *BasicQueue)  Stop() error {
	var err error
	if !bq.Started() {
		err = errors.New("Cant' stop Queue, it was not started")
	} else  if atomic.CompareAndSwapInt64(&bq.started, 1, 0) {
		err = nil
	}
	return err
}

