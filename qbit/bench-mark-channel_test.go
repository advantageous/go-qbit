package qbit

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkChannel(b *testing.B) {

	const total = 1E8
	const numRuns = 10

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

	start := time.Now().Unix()
	for run := 0; run < numRuns; run++ {

		startLoop := time.Now().Unix()

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
		counter = int32(0)

		fmt.Printf("Run %d\tElaspsed     %d seconds\n", run, time.Now().Unix()-startLoop)
		fmt.Printf("      \tTotal Time   %d seconds\n", time.Now().Unix()-start)
		memStat := &runtime.MemStats{}
		runtime.ReadMemStats(memStat)
		fmt.Printf("heap %s \t GC count %s\n", uint64ToString(memStat.HeapAlloc), uint32ToString(memStat.NumGC))

	}

}
