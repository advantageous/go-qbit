package qbit

import "time"

type BasicReceiveQueue struct {
	waitDuration    time.Duration
	channel         chan *ChannelBuffer
	lastBuffer      *ChannelBuffer
	lastBufferIndex int
}

func NewBasicReceiveQueue(waitDuration time.Duration, channel chan *ChannelBuffer) ReceiveQueue {

	return &BasicReceiveQueue{
		waitDuration: waitDuration,
		channel:      channel,
	}
}


func (brq *BasicReceiveQueue) PollWait() interface{} {

	if brq.lastBuffer != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastBuffer = brq.ReadBatchWait()
	brq.lastBufferIndex = 0
	if brq.lastBuffer != nil {
		return brq.getItemFromLocalQueue()
	} else {
		return nil
	}
}

func (brq *BasicReceiveQueue) Poll() interface{} {

	if brq.lastBuffer != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastBuffer = brq.ReadBatch()
	brq.lastBufferIndex = 0
	if brq.lastBuffer != nil {
		return brq.getItemFromLocalQueue()
	} else {
		return nil
	}

}

func (brq *BasicReceiveQueue) Take() interface{} {
	if brq.lastBuffer != nil {
		return brq.getItemFromLocalQueue()
	}
	brq.lastBuffer = brq.TakeBatch()
	brq.lastBufferIndex = 0
	return brq.getItemFromLocalQueue()
}

func (brq *BasicReceiveQueue) TakeBatch()  *ChannelBuffer {
	if brq.lastBuffer != nil {
		return brq.returnBatch()
	}
	return <-brq.channel
}

func (brq *BasicReceiveQueue) ReadBatch() *ChannelBuffer {
	if brq.lastBuffer != nil {
		return brq.returnBatch()
	}

	select {
	case items := <-brq.channel:
		return items
	default:
		return nil
	}
}

func (brq *BasicReceiveQueue) ReadBatchWait() *ChannelBuffer {
	if brq.lastBuffer != nil {
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
	item := brq.lastBuffer.Buffer[brq.lastBufferIndex]
	brq.lastBufferIndex++
	if brq.lastBufferIndex == brq.lastBuffer.Index {
		brq.lastBufferIndex = 0
		brq.lastBuffer = nil
	}
	return item
}

func (brq *BasicReceiveQueue) returnBatch() *ChannelBuffer {
	var items []interface{}

	if brq.lastBufferIndex > 0 {
		items = brq.lastBuffer.Buffer[brq.lastBufferIndex:brq.lastBuffer.Index]
	} else {
		items = brq.lastBuffer.Buffer
	}
	brq.lastBuffer = nil
	brq.lastBufferIndex = 0

	return &ChannelBuffer{items, 0}
}
