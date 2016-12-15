package qbit

import (
	"github.com/advantageous/go-qbit/logging"
//	"fmt"
)

type BasicSendQueue struct {
	outChannel     chan *ChannelBuffer
	recycleChannel chan *ChannelBuffer
	owner          Queue
	batchSize      int
	logger         logging.Logger
	buffer         *ChannelBuffer
	bufferCreate   int
	bufferRecycle   int
}

func NewSendQueue(outChannel chan *ChannelBuffer, recycleChannel chan *ChannelBuffer, owner Queue,
batchSize int, logger logging.Logger) SendQueue {

	if logger == nil {
		logger = logging.GetSimpleLogger("QBIT_SIMPLE_QUEUE", "sender")
	}

	bsq := &BasicSendQueue{
		outChannel:    outChannel,
		owner:      owner,
		batchSize:  batchSize,
		logger:     logger,
		recycleChannel: recycleChannel,
	}

	bsq.buffer = bsq.createNewBuffer()
	return bsq
}

func (bsq *BasicSendQueue) Send(item interface{}) error {
	err := bsq.flushIfOverBatch()
	if err != nil {
		return err
	}
	bsq.buffer.Buffer[bsq.buffer.Index] = item
	bsq.buffer.Index++
	return err
}

func (bsq *BasicSendQueue) flushIfOverBatch() error {
	if bsq.buffer.Index < bsq.batchSize {
		return nil
	} else {
		return bsq.sendLocalQueue()
	}
}

func (bsq *BasicSendQueue) sendLocalQueue() error {
	if bsq.buffer.Index > 0 {
		bsq.outChannel <- bsq.buffer
		bsq.buffer = bsq.createBuffer()
	}
	return nil
}

func (bsq *BasicSendQueue) createNewBuffer() *ChannelBuffer {
	return &ChannelBuffer{
		Buffer: make([]interface{}, bsq.batchSize),
		Index: 0,
	}
}

func (bsq *BasicSendQueue) createBuffer() *ChannelBuffer {

	var newBuffer *ChannelBuffer

	select {
	case buf := <-bsq.recycleChannel:
		buf.Index = 0
		newBuffer = buf
		bsq.bufferRecycle++
		//if bsq.bufferRecycle % 10000 == 0 {
		//	fmt.Println("RECYCLE BUF", bsq.bufferRecycle)
		//}

	default:
		newBuffer = bsq.createNewBuffer()
		bsq.bufferCreate++
		//if bsq.bufferCreate % 10 == 0 {
		//	fmt.Println("CREATED BUF", bsq.bufferCreate)
		//}
	}

	return newBuffer
}

func (bsq *BasicSendQueue) FlushSends() error {
	return bsq.sendLocalQueue()
}

func (bsq *BasicSendQueue) Size() int {
	return len(bsq.outChannel)
}
