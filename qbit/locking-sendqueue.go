package qbit

import (
	"sync"
	"time"
)

type LockingSendQueue struct {
	sendQueue SendQueue
	rw        sync.RWMutex
}

func NewLockingSendQueue(sendQueue SendQueue) SendQueue {
	return &LockingSendQueue{
		rw:        sync.RWMutex{},
		sendQueue: sendQueue,
	}
}

func NewLockingSendQueueWithAutoFlush(sendQueue SendQueue, flushDuration time.Duration) SendQueue {

	newSendQueue := NewLockingSendQueue(sendQueue)

	go func() {
		for {
			timer := time.NewTimer(flushDuration)
			select {
			case <-timer.C:
				newSendQueue.FlushSends()
				timer.Reset(flushDuration)
			}
		}
	}()

	return newSendQueue
}

func (lq *LockingSendQueue) Send(item interface{}) error {
	lq.rw.Lock()
	defer lq.rw.Unlock()
	return lq.sendQueue.Send(item)
}

func (lq *LockingSendQueue) FlushSends() error {
	lq.rw.Lock()
	defer lq.rw.Unlock()
	return lq.sendQueue.FlushSends()
}

func (lq *LockingSendQueue) Size() int {
	lq.rw.RLock()
	defer lq.rw.RUnlock()
	return lq.sendQueue.Size()
}
