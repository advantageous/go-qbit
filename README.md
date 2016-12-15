# go-qbit
GoLang QBit, just the queue!

Reduces thread synchronization and locking by sending messages in batches. 
Implements auto-flushing of batches. 


QBit queue is over 7.2 to 7.5 times faster than golang's internal channel.
You also get notification via callbacks when the queue is empty, a new batch started,
or a new batch ended (and more).


#### QBit queue
```sh
go test -v github.com/advantageous/go-qbit/qbit -bench ^BenchmarkQueue$ -run ^$
1	13982813803 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	13.992s

```

QBit with a single reader and a single writer was able to send 
71,469,411 messages a second. 

#### Go Channel 
```sh
go test -v github.com/advantageous/go-qbit/qbit -bench ^BenchmarkChannel$ -run ^$
1	104966302989 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	104.980s

```

Golang channel with a single reader and a single writer was able to send 
9,525,623 messages a second. 


#### Improvement 

Use pooling. Created simpler queue manager. 

#### Queue Pooling off (2 seconds faster than before with simpler queue manager)
```
heap 2,437,776 	 GC count 4,531
1	12666113306 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	12.676s
```


#### Queue Pooling on (3 seconds faster than before with simpler queue manager)
```
heap 2,505,232 	 GC count 3
1	10912673726 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	10.924s
```

GC collections went from 4,531 to 3. 

![perf](/images/perf.png)


