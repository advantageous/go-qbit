package qbit

func manageQueue(limit int, queue Queue, inputQueue ReceiveQueue, listener ReceiveQueueListener) {
	listener.Init()
	var item interface{}
	count := 0
	item = inputQueue.Poll() //Initialize things.

	OuterLoop:
	for {
		if (item != nil) {
			listener.StartBatch()
		}

		for {
			if item == nil {
				break
			}
			listener.Receive(item)
			/* If the receive count has hit the max then we need to call limit. */
			if (count >= limit) {
				listener.Limit()
				count = 0
				if queue.Stopped() {
					listener.Shutdown()
					break OuterLoop
				}
			}
			/* Grab the next item from the queue. */
			item = inputQueue.Poll()
			count++
		}

		count = 0
		listener.Empty()

		// Get the next item, but wait this time since the queue was empty.
		// This pauses the queue handling so we don't eat up all of the CPU.
		item = inputQueue.PollWait()
		if queue.Stopped() {
			listener.Shutdown()
			break OuterLoop
		}

		if item == nil {
			/* Idle means we yielded and then waited a full wait time, so idle might be a good time to do clean up
			or timed tasks.
			 */
			listener.Idle();
		}
	}
}





