package qbit

import "sync"

var pools map[int]*sync.Pool = make(map[int]*sync.Pool)
var rw = sync.RWMutex{}

const poolOn = true

func readPool(batchSize int) (*sync.Pool, bool) {
	pool, ok := pools[batchSize]
	return pool, ok
}

func getPool(batchSize int) *sync.Pool {
	rw.Lock()
	defer rw.Unlock()

	pool, ok := readPool(batchSize)
	if ok {
		return pool
	} else {
		return createPool(batchSize)
	}
}

func createPool(batchSize int) *sync.Pool {
	pool := &sync.Pool{}
	pool.New = func() interface{} {
		return make([]interface{}, batchSize)
	}
	pools[batchSize] = pool
	return pool
}

func makeBuffer(batchSize int) []interface{} {
	if poolOn {
		return getPool(batchSize).Get().([]interface{})
	} else {
		return make([]interface{}, batchSize)
	}
}

func recycleBuffer(batchSize int, buffer []interface{}) {
	if poolOn {
		getPool(batchSize).Put(buffer)
	}
}
