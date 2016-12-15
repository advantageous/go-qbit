package qbit

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
	//	"runtime"
	"runtime"
	"strconv"
)

func BenchmarkQueue(b *testing.B) {

	const total = 1E8
	const numRuns = 10

	counter := int32(0)

	type Send struct {
		name string
	}

	var testItem *Send

	testItem = &Send{
		name: "Foo",
	}

	queueManager := NewSimpleQueueManager(NewReceiveListener(func(interface{}) {
		atomic.AddInt32(&counter, 1)

	}))

	b.ResetTimer()

	start := time.Now().Unix()
	for run := 0; run < numRuns; run++ {

		startLoop := time.Now().Unix()

		go func() {
			sendQueue := queueManager.Queue().SendQueue()
			for i := 0; i < total; i++ {
				sendQueue.Send(testItem)
			}
			sendQueue.FlushSends()
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
		counter = int32(0)
	}

}

func uint64ToString(n uint64) string {
	in := strconv.FormatUint(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func uint32ToString(n uint32) string {
	in := strconv.FormatUint(uint64(n), 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
