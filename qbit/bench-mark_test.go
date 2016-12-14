package qbit

import (
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkQueue(b *testing.B) {

	const total = 1E8

	counter := int32(0)

	queueManager := NewQueueManager(10, 10000, 10000, 10*time.Millisecond, NewReceiveListener(func(interface{}) {
		atomic.AddInt32(&counter, 1)

	}))
	b.ResetTimer()

	sendQueue := queueManager.Queue().SendQueue()

	go func() {
		for i := 0; i < total; i++ {
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
		<-time.NewTimer(100 * time.Millisecond).C

	}()

	for {
		<-time.NewTimer(10 * time.Millisecond).C
		count := atomic.LoadInt32(&counter)
		if count >= total {
			break
		}
	}

}
