package qbit

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkQueue(b *testing.B) {

	const total = 10E7

	counter := int32(0)

	queueManager := NewQueueManager(3, 1500, 1500, 1*time.Millisecond, NewReceiveListener(func(interface{}) {
		atomic.AddInt32(&counter, 1)
	}))

	sendQueue := queueManager.Queue().SendQueue()
	b.ResetTimer()

	for runs:=0; runs < 1; runs++ {

		go func() {
			for i := 0; i < total; i++ {
				sendQueue.Send(i)
			}
			sendQueue.FlushSends()
		}()

		for {
			count := atomic.LoadInt32(&counter)
			if count >= total {
				atomic.StoreInt32(&counter, 0)
				break
			} else {
				<-time.NewTimer(10 * time.Millisecond).C
			}

		}
	}

}
