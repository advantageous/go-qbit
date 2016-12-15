package qbit

import (
	"testing"
	"time"
	"sync/atomic"
)

func BenchmarkChannel(b *testing.B) {

	const total = 1E8 // 100,000,000

	channel := make(chan int, 1000)
	counter := int32(0)

	b.ResetTimer()

	go func() {
		for i := 0; i < total; i++ {
			channel <- i
		}

	}()

	go func() {
		for {
			<-channel
			atomic.AddInt32(&counter, 1)
		}
	}()

	for {
		<-time.NewTimer(10 * time.Millisecond).C
		count := atomic.LoadInt32(&counter)
		if count >= total {
			break
		}
	}


}
