package qbit

import (
	lg "github.com/advantageous/go-qbit/logging"
	tlg "github.com/advantageous/go-qbit/logging/test"
	"testing"
	"time"
)

func TestBasics(t *testing.T) {
	logger := tlg.NewTestDebugLogger("test", t)

	channel := make(chan []interface{}, 10)

	queueReceiver := NewBasicReceiveQueue(time.Millisecond*10, channel)

	channel <- []interface{}{"Hello"}
	channel <- []interface{}{"How", "Are", "You"}

	var item string
	var ok bool
	item, ok = queueReceiver.Poll().(string)

	if item != "Hello" && ok {
		logger.Error("Item not Hello", item)
	}

	item = queueReceiver.Take().(string)

	if item != "How" {
		logger.Error("Item not How", item)
	}

	item = queueReceiver.PollWait().(string)

	if item != "Are" {
		logger.Error("Item not Are", item)
	}

	items := queueReceiver.ReadBatch()

	if len(items) <= 0 {
		logger.Error("Items wrong length", len(items))
	}

	item = items[0].(string)

	if item != "You" {
		logger.Error("Item not You", item)
	}
}

func setupSinglePoll(t *testing.T) (ReceiveQueue, lg.Logger) {
	logger := tlg.NewTestDebugLogger("test", t)
	channel := make(chan []interface{}, 10)
	queueReceiver := NewBasicReceiveQueue(time.Millisecond*10, channel)

	channel <- []interface{}{"0a", "0b", "0c"}
	channel <- []interface{}{"1a", "1b", "1c"}
	channel <- []interface{}{"2a", "2b", "2c"}
	return queueReceiver, logger
}

func setup(t *testing.T) (ReceiveQueue, lg.Logger) {
	logger := tlg.NewTestDebugLogger("test", t)

	channel := make(chan []interface{}, 10)

	queueReceiver := NewBasicReceiveQueue(time.Millisecond*10, channel)

	channel <- []interface{}{"How", "Are", "You"}
	return queueReceiver, logger

}

func TestPoll(t *testing.T) {
	queueReceiver, logger := setupSinglePoll(t)
	pollMethod := queueReceiver.Poll
	testPoll(pollMethod, logger, queueReceiver)
}

func TestTake(t *testing.T) {
	queueReceiver, logger := setupSinglePoll(t)
	pollMethod := queueReceiver.Take
	testPoll(pollMethod, logger, queueReceiver)
}

func TestPollWait(t *testing.T) {
	queueReceiver, logger := setupSinglePoll(t)
	pollMethod := queueReceiver.PollWait
	testPoll(pollMethod, logger, queueReceiver)
}

func TestEmpty(t *testing.T) {
	logger := tlg.NewTestDebugLogger("test", t)
	channel := make(chan []interface{}, 10)
	queueReceiver := NewBasicReceiveQueue(time.Millisecond*500, channel)

	item := queueReceiver.PollWait()

	if item != nil {
		logger.Error("Items should be nil")
	}

	item = queueReceiver.Poll()
	if item != nil {
		logger.Error("Items should be nil")
	}

	items := queueReceiver.ReadBatch()
	if items != nil {
		logger.Error("Items should be nil")
	}

	items = queueReceiver.ReadBatchWait()
	if items != nil {
		logger.Error("Items should be nil")
	}

}

func testPoll(pollMethod func() interface{}, logger lg.Logger, receiveQueue ReceiveQueue) {

	var item string
	var ok bool

	if receiveQueue.Size() != 3 {
		logger.Error("Wrong size", receiveQueue.Size())
	}

	for i := 0; i < 9; i++ {
		item, ok = pollMethod().(string)
		if !ok {
			logger.Error("Error cast")
		}
		if len(item) != 2 {
			logger.Error("Error size", item)
		}
	}

	if receiveQueue.Size() != 0 {
		logger.Error("Wrong size After", receiveQueue.Size())
	}
}

func TestReadBatchWholeBatch(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.ReadBatch
	testReadBatchWhole(readBatch, logger)
}

func TestTakeBatchWholeBatch(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.TakeBatch
	testReadBatchWhole(readBatch, logger)
}

func TestReadBatchWholeBatchWait(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.ReadBatchWait
	testReadBatchWhole(readBatch, logger)
}

func TestReadBatchPartBatch(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.ReadBatch
	queueReceiver.Take()
	testReadBatchWholePart(readBatch, logger)
}

func TestReadBatchPartBatchWait(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.ReadBatchWait
	queueReceiver.Take()
	testReadBatchWholePart(readBatch, logger)
}

func TestTakeBatchPartBatc(t *testing.T) {
	queueReceiver, logger := setup(t)
	readBatch := queueReceiver.TakeBatch
	queueReceiver.Take()
	testReadBatchWholePart(readBatch, logger)
}

func testReadBatchWhole(readBatch func() []interface{}, logger lg.Logger) {
	items := readBatch()

	if len(items) != 3 {
		logger.Error("Items wrong length", len(items))
	}

	item := items[0].(string)

	if item != "How" {
		logger.Error("Item not How", item)
	}

}

func testReadBatchWholePart(readBatch func() []interface{}, logger lg.Logger) {
	items := readBatch()

	if len(items) != 2 {
		logger.Error("Items wrong length", len(items))
	}

	item := items[0].(string)

	if item != "Are" {
		logger.Error("Item not Are", item)
	}

}
