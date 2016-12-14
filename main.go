package main

import (
	lg "github.com/advantageous/go-qbit/logging"
	q "github.com/advantageous/go-qbit/qbit"
	"strconv"
	"time"
)

func main() {

	logger := lg.NewSimpleLogger("main")

	queue := q.NewActiveQueue(10, 10, 10, "test", time.Millisecond * 100)
	sendQueue := queue.SendQueue()
	channel := make(chan interface{}, 100)

	listener := q.NewListenerReceive(func(item interface{}) {
		channel <- item
	})

	addItems := func() {
		for i := 0; i < 100; i++ {
			timer := time.NewTimer(100 * time.Millisecond)

			select {
			case <-timer.C:

			}
			sendQueue.Send(strconv.Itoa(i))
		}
		sendQueue.FlushSends()
	}

	go addItems()

	queue.StartListener(listener)

	drainItems := func() {
		count := 0
		for item := range channel {
			logger.Debug(item)
			count++
			if count == 100 {
				break
			}
		}

	}

	go drainItems()

	for {
		queue.Stop()
	}

}
