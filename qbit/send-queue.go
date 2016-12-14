package qbit

import (
	"github.com/advantageous/go-qbit/logging"
)

type BasicSendQueue struct {
	channel    chan []interface{}
	owner      Queue
	batchSize  int
	logger     logging.Logger
	index      int
	queueLocal []interface{}
}

func NewSendQueue(channel chan []interface{}, owner Queue, batchSize  int, logger logging.Logger) SendQueue{

	if logger == nil {
		logger = logging.GetSimpleLogger("QBIT_SIMPLE_QUEUE", owner.Name() + "-sender")
	}

	queueLocal := make([]interface{}, batchSize)

	return &BasicSendQueue{
		channel: channel,
		owner: owner,
		batchSize: batchSize,
		logger: logger,
		queueLocal: queueLocal,
	}
}

func (bsq *BasicSendQueue) Send(item interface{}) (bool, error) {

	ableToSend := bsq.flushIfOverBatch()
	bsq.queueLocal[bsq.index] = item
	bsq.index++;
	return ableToSend, nil;
}

func (bsq *BasicSendQueue) flushIfOverBatch() bool {
	return bsq.index < bsq.batchSize || bsq.sendLocalQueue()
}

func (bsq *BasicSendQueue) sendLocalQueue() bool {
	if bsq.index > 0 {
		slice := make([]interface{}, bsq.index)
		copy(slice, bsq.queueLocal)

		select {
		case bsq.channel <- slice:
			bsq.index = 0
			for i := 0; i < len(bsq.queueLocal); i++ {
				bsq.queueLocal[i] = nil
			}
			return true
		default:
			return false
		}
	} else {
		return true;
	}
}

func (bsq *BasicSendQueue) FlushSends() error {
	bsq.sendLocalQueue()
	return nil
}

func (bsq *BasicSendQueue) Size() int {
	return len(bsq.channel)
}

func (bsq *BasicSendQueue) Name() string {
	return bsq.owner.Name()
}
