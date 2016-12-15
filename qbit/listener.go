package qbit

func NewListener() ReceiveQueueListener {
	return &BaseReceiveQueueListener{}
}

func NewReceiveListener(receiveFunc func(item interface{})) ReceiveQueueListener {
	return NewQueueListener(&QueueListener{ReceiveFunc: receiveFunc})
}

var EmptyFunc func() = func() {}

func NewQueueListener(queueListener *QueueListener) ReceiveQueueListener {

	if queueListener.InitFunc == nil {
		queueListener.InitFunc = EmptyFunc
	}
	if queueListener.ReceiveFunc == nil {
		queueListener.ReceiveFunc = func(item interface{}) {}
	}
	if queueListener.EmptyFunc == nil {
		queueListener.EmptyFunc = EmptyFunc
	}
	if queueListener.ShutdownFunc == nil {
		queueListener.ShutdownFunc = EmptyFunc
	}
	if queueListener.IdleFunc == nil {
		queueListener.IdleFunc = EmptyFunc
	}
	if queueListener.StartBatchFunc == nil {
		queueListener.StartBatchFunc = EmptyFunc
	}
	if queueListener.EndBatchFunc == nil {
		queueListener.EndBatchFunc = EmptyFunc
	}
	return &BaseReceiveQueueListener{
		init:       queueListener.InitFunc,
		receive:    queueListener.ReceiveFunc,
		empty:      queueListener.EmptyFunc,
		shutdown:   queueListener.ShutdownFunc,
		idle:       queueListener.IdleFunc,
		startBatch: queueListener.StartBatchFunc,
		limit:      queueListener.EndBatchFunc,
	}
}

type QueueListener struct {
	InitFunc       func()
	ReceiveFunc    func(item interface{})
	EmptyFunc      func()
	ShutdownFunc   func()
	IdleFunc       func()
	StartBatchFunc func()
	EndBatchFunc      func()
}

type BaseReceiveQueueListener struct {
	init       func()
	receive    func(item interface{})
	empty      func()
	shutdown   func()
	idle       func()
	startBatch func()
	limit      func()
}

func (l *BaseReceiveQueueListener) EndBatch() {
	l.limit()
}
func (l *BaseReceiveQueueListener) Init() {
	l.init()
}
func (l *BaseReceiveQueueListener) Receive(item interface{}) {
	l.receive(item)
}
func (l *BaseReceiveQueueListener) Empty() {
	l.empty()
}
func (l *BaseReceiveQueueListener) Shutdown() {
	l.shutdown()
}
func (l *BaseReceiveQueueListener) Idle() {
	l.idle()
}
func (l *BaseReceiveQueueListener) StartBatch() {
	l.startBatch()
}
