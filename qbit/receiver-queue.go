package qbit

import "time"

type BasicReceiveQueue struct {
	waitDuration   time.Duration
	channel        chan []interface{}
	lastQueue      []interface{}
	lastQueueIndex int
}

func NewReceiveQueue(waitDuration time.Duration, channel chan []interface{}) ReceiveQueue {

	return &BasicReceiveQueue{
		waitDuration: waitDuration,
		channel:      channel,
	}
}

func NewSimpleReceiveQueue(channel chan []interface{}) ReceiveQueue {

	return NewReceiveQueue(time.Millisecond*10, channel)
}

func (brq *BasicReceiveQueue) PollWait() interface{} {

	if brq.lastQueue != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastQueue = brq.ReadBatchWait()
	brq.lastQueueIndex = 0
	if brq.lastQueue != nil {
		return brq.getItemFromLocalQueue()
	} else {
		return nil
	}
}

func (brq *BasicReceiveQueue) Poll() interface{} {

	if brq.lastQueue != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastQueue = brq.ReadBatch()
	brq.lastQueueIndex = 0
	if brq.lastQueue != nil {
		return brq.getItemFromLocalQueue()
	} else {
		return nil
	}

}

func (brq *BasicReceiveQueue) Take() interface{} {
	if brq.lastQueue != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastQueue = brq.TakeBatch()
	brq.lastQueueIndex = 0
	return brq.getItemFromLocalQueue()
}

func (brq *BasicReceiveQueue) TakeBatch() []interface{} {
	if brq.lastQueue != nil {
		return brq.returnBatch()
	}
	return <-brq.channel
}

func (brq *BasicReceiveQueue) ReadBatch() []interface{} {
	if brq.lastQueue != nil {
		return brq.returnBatch()
	}

	select {
	case items := <-brq.channel:
		return items
	default:
		return nil
	}
}

func (brq *BasicReceiveQueue) ReadBatchWait() []interface{} {
	if brq.lastQueue != nil {
		return brq.returnBatch()
	}

	timer := time.NewTimer(brq.waitDuration)

	select {
	case items := <-brq.channel:
		timer.Stop()
		return items
	case <-timer.C:
		return nil
	}
}

func (brq *BasicReceiveQueue) Size() int {
	return len(brq.channel)
}

func (brq *BasicReceiveQueue) getItemFromLocalQueue() interface{} {
	item := brq.lastQueue[brq.lastQueueIndex]
	brq.lastQueueIndex++
	if brq.lastQueueIndex == len(brq.lastQueue) {
		brq.lastQueueIndex = 0
		brq.lastQueue = nil
	}
	return item
}

func (brq *BasicReceiveQueue) returnBatch() []interface{} {
	var items []interface{}

	if brq.lastQueueIndex > 0 {
		items = brq.lastQueue[brq.lastQueueIndex:len(brq.lastQueue)]
	} else {
		items = brq.lastQueue
	}
	brq.lastQueue = nil
	brq.lastQueueIndex = 0
	return items
}
