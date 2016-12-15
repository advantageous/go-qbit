package main

import (
	"time"
	"sync/atomic"
	q "github.com/advantageous/go-qbit/qbit"
	"strconv"
	"fmt"
)


func main() {

	const total = 1000000

	counter := int32(0)

	queueManager := q.NewQueueManager(1000, 10, 10, 10*time.Millisecond, q.NewReceiveListener(func(interface{}) {
		atomic.AddInt32(&counter, 1)

		count := atomic.LoadInt32(&counter)
		if count % 100 == 0{
			println("COUNT FROM REC", count)
		}
	}))

	sendQueue := queueManager.Queue().SendQueue()

	go func() {
		for i := 0; i < total; i++ {
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
		<-time.NewTimer(100 * time.Millisecond).C

	}()

	for {
		<-time.NewTimer(1000 * time.Millisecond).C
		count := atomic.LoadInt32(&counter)
		if count > 100 {
			fmt.Println("COUNT", count)
		}
		if count > 1000 {
			fmt.Println("COUNT", count)
		}
		if count > 10000 {
			fmt.Println("COUNT", count)
		}
		if count >= total {
			break
		}
	}

}
