package qbit

import (
	"errors"
	"sync/atomic"
	"time"
)

type BasicQueueManager struct {
	queue     Queue
	started   int64
	batchSize int
}

func NewQueueManager(channelSize int, batchSize int, pollWaitDuration time.Duration, listener ReceiveQueueListener) QueueManager {
	channel := make(chan []interface{}, channelSize)
	queue := &BasicQueue{
		channel:          channel,
		batchSize:        batchSize,
		pollWaitDuration: pollWaitDuration,
	}

	queueManager := &BasicQueueManager{
		queue:     queue,
		batchSize: batchSize,
	}

	if listener == nil {
		listener = NewQueueListener(&QueueListener{})
	}
	queueManager.startListener(listener)
	return queueManager
}

func NewSimpleQueueManager(listener ReceiveQueueListener) QueueManager {
	return NewQueueManager(1, 10000, time.Millisecond*10, listener)
}

func (bqm *BasicQueueManager) startListener(listener ReceiveQueueListener) error {
	var err error

	if bqm.Started() {
		err = errors.New("Queue already started")
	} else if atomic.CompareAndSwapInt64(&bqm.started, 0, 1) {
		go manageQueue(bqm.batchSize, bqm, bqm.queue.ReceiveQueue(), listener)
	}
	return err
}

func (bqm *BasicQueueManager) Started() bool {
	started := atomic.LoadInt64(&bqm.started)
	return started == 1
}

func (bqm *BasicQueueManager) Stopped() bool {
	started := atomic.LoadInt64(&bqm.started)
	return started == 0
}

func (bqm *BasicQueueManager) Queue() Queue {
	return bqm.queue
}

func (bqm *BasicQueueManager) SendQueueWithAutoFlush(flushDuration time.Duration) SendQueue {

	sendQueue := NewLockingSendQueue(bqm.queue.SendQueue())

	go func() {
		for {
			timer := time.NewTimer(flushDuration)
			select {
			case <-timer.C:
				sendQueue.FlushSends()
				timer.Reset(flushDuration)
			}

			if bqm.Stopped() {
				break
			}
		}
	}()

	return sendQueue
}

func (bqm *BasicQueueManager) Stop() error {
	var err error
	if !bqm.Started() {
		err = errors.New("Cant' stop Queue, it was not started")
	} else if atomic.CompareAndSwapInt64(&bqm.started, 1, 0) {
		err = nil
	}
	return err
}

func manageQueue(batchSize int, queueManager QueueManager, inputQueue ReceiveQueue,
	listener ReceiveQueueListener) {
	listener.Init()
	var items []interface{}
	count := 0

	items = inputQueue.ReadBatch()

OuterLoop:
	for {

		if items != nil {
			listener.StartBatch()
			for i := 0; i < len(items); i++ {
				count++
				listener.Receive(items[i])
			}
			listener.EndBatch()
			if batchSize == len(items) {
				recycleBuffer(batchSize, items)
			}
			items = inputQueue.ReadBatch()
			continue OuterLoop
		} else {
			listener.Empty()
			items = inputQueue.ReadBatchWait()
			if items == nil {
				listener.Idle()
				if queueManager.Stopped() {
					listener.Shutdown()
					break OuterLoop
				}
			}
		}
	}
}

