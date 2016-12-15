package qbit

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkChannel(b *testing.B) {

	const total = 1E9
	type Send struct {
		name string
	}

	var testItem *Send

	testItem = &Send{
		name: "Foo",
	}

	channel := make(chan *Send, 1000)
	counter := int32(0)

	b.ResetTimer()

	go func() {
		for i := 0; i < total; i++ {
			channel <- testItem
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
