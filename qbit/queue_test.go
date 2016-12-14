package qbit

import (
	tlg "github.com/advantageous/go-qbit/logging/test"
	"strconv"
	"testing"
	"time"
)

func TestQueueBasics(t *testing.T) {
	logger := tlg.NewTestDebugLogger("test", t)

	queue := NewQueue(10, 10, 10, time.Millisecond*100)
	queueReceiver := queue.ReceiveQueue()
	sendQueue := queue.SendQueue()

	sendQueue.Send("Hello")
	sendQueue.Send("How")
	sendQueue.Send("Are")
	sendQueue.Send("You")
	sendQueue.FlushSends()

	if sendQueue.Size() != 1 {
		logger.Error("Size is wrong")
	}

	if queue.Size() != 1 {
		logger.Error("Size is wrong")
	}

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

func TestOverBatch(t *testing.T) {
	logger := tlg.NewTestSimpleLogger("test", t)

	queue := NewQueue(5, 100, 10, time.Millisecond*100)
	queueReceiver := queue.ReceiveQueue()
	sendQueue := queue.SendQueue()

	for i := 0; i < 100; i++ {
		sendQueue.Send(strconv.Itoa(i))
	}

	if sendQueue.Size() != 19 {
		logger.Error("Size is wrong", sendQueue.Size())
	}

	if queue.Size() != 19 {
		logger.Error("Size is wrong")
	}
	sendQueue.FlushSends()

	count := 0
	for ; count < 100; count++ {
		logger.Debug(queueReceiver.Poll().(string))
	}

	if count != 100 {
		logger.Error("Count is wrong", count)
	}
}

func TestQueueAsync(t *testing.T) {
	logger := tlg.NewTestSimpleLogger("test", t)

	queue := NewQueue(5, 100, 10, time.Millisecond*100)
	sendQueue := queue.SendQueue()
	channel := make(chan interface{})

	listener := NewListenerReceive(func(item interface{}) {
		channel <- item
	})

	addItems := func() {
		for i := 0; i < 100; i++ {
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
	}

	go addItems()

	queue.StartListener(listener)

	count := 0
	for item := range channel {
		logger.Debug(item)
		count++
		if count == 100 {
			break
		}
	}

	if count != 100 {
		logger.Error("Count should be 100")
	}
	if len(channel) != 0 {
		logger.Error("Channel should be empty", len(channel))
	}
}

func TestQueueAsyncStop(t *testing.T) {

	logger := tlg.NewTestSimpleLogger("test", t)

	queue := NewQueue(5, 100, 10, time.Millisecond*100)
	sendQueue := queue.SendQueue()
	channel := make(chan interface{})

	listener := NewListenerReceive(func(item interface{}) {
		channel <- item
	})

	addItems := func() {
		for i := 0; i < 100; i++ {

			timer := time.NewTimer(1 * time.Second)
			select {
			case <-timer.C:
				timer.Stop()
			}

			if queue.Stopped() {
				break
			}
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
	}

	go addItems()

	queue.StartListener(listener)
	queue.StartListener(listener)
	queue.Stop()
	queue.Stop()

	if len(channel) > 0 {
		logger.Error("Channel should be not empty", len(channel))
	}
}
