package qbit

import (
	"sync"
)

type LockingSendQueue struct {
	sendQueue SendQueue
	rw sync.RWMutex
}


func NewLockingSendQueue(sendQueue SendQueue) SendQueue {
	rw := sync.RWMutex{}
	return &LockingSendQueue{
		rw : rw,
		sendQueue: sendQueue,
	}
}

func (lq *LockingSendQueue) Send(item interface{})  error  {
	lq.rw.Lock()
	defer lq.rw.Unlock()
	return lq.sendQueue.Send(item)
}

func (lq *LockingSendQueue) FlushSends()  error {
	lq.rw.Lock()
	defer lq.rw.Unlock()
	return lq.sendQueue.FlushSends()
}

func (lq *LockingSendQueue) Name()  string {
	return lq.sendQueue.Name()
}

func (lq *LockingSendQueue) Size()  int  {
	lq.rw.RLock()
	defer lq.rw.RUnlock()
	return lq.sendQueue.Size()
}

