package qbit

import (
	"github.com/advantageous/go-logback/logging"
)

type BasicSendQueue struct {
	channel   chan []interface{}
	owner     Queue
	batchSize int
	logger    logging.Logger
	index     int
	buffer    []interface{}
}

func NewSendQueue(channel chan []interface{}, owner Queue, batchSize int, logger logging.Logger) SendQueue {

	if logger == nil {
		logger = logging.GetSimpleLogger("QBIT_SIMPLE_QUEUE", "sender")
	}

	buffer := makeBuffer(batchSize)

	return &BasicSendQueue{
		channel:   channel,
		owner:     owner,
		batchSize: batchSize,
		logger:    logger,
		buffer:    buffer,
	}
}

func (bsq *BasicSendQueue) Send(item interface{}) error {
	err := bsq.flushIfOverBatch()
	if err != nil {
		return err
	}
	bsq.buffer[bsq.index] = item
	bsq.index++
	return err
}

func (bsq *BasicSendQueue) flushIfOverBatch() error {
	if bsq.index < bsq.batchSize {
		return nil
	} else {
		return bsq.sendLocalQueue()
	}
}

func (bsq *BasicSendQueue) sendLocalQueue() error {
	if bsq.index > 0 {
		bsq.channel <- bsq.buffer[0:bsq.index]
		bsq.index = 0
		bsq.buffer = makeBuffer(bsq.batchSize)
	}
	return nil
}

func (bsq *BasicSendQueue) FlushSends() error {
	return bsq.sendLocalQueue()
}

func (bsq *BasicSendQueue) Size() int {
	return len(bsq.channel)
}
