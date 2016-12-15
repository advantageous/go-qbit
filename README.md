# go-qbit
GoLang QBit, just the queue

Reduces thread synchronization and locking by sending messages in batches. 
Implements auto-flushing of batches. 


QBit queue is over 7.2 to 7.5 times faster than golang's internal channel.
You also get notification via callbacks when the queue is empty, a new batch started,
or a new batch ended. 


#### QBit queue
```sh
go test -v github.com/advantageous/go-qbit/qbit -bench ^BenchmarkQueue$ -run ^$
1	13982813803 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	13.992s

```

#### Go Channel 
```sh
go test -v github.com/advantageous/go-qbit/qbit -bench ^BenchmarkChannel$ -run ^$
1	104966302989 ns/op
PASS
ok  	github.com/advantageous/go-qbit/qbit	104.980s

```


