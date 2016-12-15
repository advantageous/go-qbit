package qbit

import (
	tlg "github.com/advantageous/go-qbit/logging/test"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestQueueBasics(t *testing.T) {
	logger := tlg.NewTestDebugLogger("test", t)

	listener := NewQueueListener(&QueueListener{})
	queue := NewQueueManager(10, 10, time.Millisecond*100, listener)
	queueReceiver := queue.Queue().ReceiveQueue()
	sendQueue := queue.Queue().SendQueue()

	sendQueue.Send("Hello")
	sendQueue.Send("How")
	sendQueue.Send("Are")
	sendQueue.Send("You")
	sendQueue.FlushSends()

	if sendQueue.Size() != 1 {
		logger.Error("Size is wrong")
	}

	if queue.Queue().Size() != 1 {
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

	queue := NewQueue(5, 100, time.Millisecond*100)
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

	channel := make(chan interface{})

	listener := NewReceiveListener(func(item interface{}) {
		channel <- item
	})

	queue := NewQueueManager(5, 100, time.Millisecond*100, listener)
	sendQueue := queue.Queue().SendQueue()

	addItems := func() {
		for i := 0; i < 100; i++ {
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
	}

	go addItems()

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

func TestListener(t *testing.T) {
	logger := tlg.NewTestSimpleLogger("test", t)
	channel := make(chan interface{})

	var initCalled, shutdownCalled, limitCalled, idleCalled, startBatchCalled, emptyCalled int64

	listener := NewQueueListener(&QueueListener{
		InitFunc: func() {
			atomic.AddInt64(&initCalled, 1)
		},
		ShutdownFunc: func() {
			atomic.AddInt64(&shutdownCalled, 1)
		},
		LimitFunc: func() {
			atomic.AddInt64(&limitCalled, 1)
		},
		IdleFunc: func() {
			atomic.AddInt64(&idleCalled, 1)
		},
		StartBatchFunc: func() {
			atomic.AddInt64(&startBatchCalled, 1)
		},
		EmptyFunc: func() {
			atomic.AddInt64(&emptyCalled, 1)
		},
		ReceiveFunc: func(item interface{}) {
			channel <- item
		},
	})

	queue := NewQueueManager(5, 100, time.Millisecond*20, listener)
	sendQueue := queue.SendQueueWithAutoFlush(10 * time.Millisecond)

	addItems := func() {
		for i := 0; i < 100; i++ {
			sendQueue.Send(strconv.Itoa(i))

			if i > 70 && i%10 == 0 {
				timer := time.NewTimer(100 * time.Millisecond)
				<-timer.C
			}
		}
		sendQueue.FlushSends()
	}

	go addItems()

	count := 0
	for item := range channel {
		logger.Debug(item)
		count++
		if count == 100 {
			break
		}
	}

	logger.Info("AFTER LOOP")

	<-time.NewTimer(100 * time.Millisecond).C

	if count != 100 {
		logger.Error("Count should be 100")
	}
	if len(channel) != 0 {
		logger.Error("Channel should be empty", len(channel))
	}

	queue.Stop()

	<-time.NewTimer(100 * time.Millisecond).C

	logger.Infof("\ninitCalled %d, shutdownCalled %d, limitCalled %d, "+
		"\nidleCalled %d, startBatchCalled %d, emptyCalled %d",
		initCalled, shutdownCalled, limitCalled,
		idleCalled, startBatchCalled, emptyCalled)

	NewQueueListener(&QueueListener{})

}

func TestQueueAsyncStop(t *testing.T) {

	logger := tlg.NewTestSimpleLogger("test", t)
	channel := make(chan interface{})

	listener := NewReceiveListener(func(item interface{}) {
		channel <- item
	})

	queue := NewQueueManager(5, 100, time.Millisecond*100, listener)
	sendQueue := queue.Queue().SendQueue()

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

	queue.Stop()
	queue.Stop()

	if len(channel) > 0 {
		logger.Error("Channel should be not empty", len(channel))
	}
}

func TestAutoFlush(t *testing.T) {

	logger := tlg.NewTestSimpleLogger("test", t)

	queue := NewQueue(5, 100, time.Millisecond*100)
	sendQueue := NewLockingSendQueueWithAutoFlush(queue.SendQueue(), time.Millisecond*100)
	receiveQueue := queue.ReceiveQueue()

	sendQueue.Send("Hi Mom")

	item := receiveQueue.Take()

	if item != "Hi Mom" {
		logger.Error("Item is not equal to Hi Mom")
	}

	if sendQueue.Size() != 0 {
		logger.Error("Size wrong")
	}

}
