package qbit

<<<<<<< HEAD
func NewListener() ReceiveQueueListener {
	return &BaseReceiveQueueListener{}
}

func NewListenerReceive(receiveFunc func(item interface{})) ReceiveQueueListener {

	return &BaseReceiveQueueListener{receive: receiveFunc}
=======
func NewListenerReceive(receiveFunc func(item interface{})) ReceiveQueueListener {
	return NewQueueListener(&QueueListener{Receive:receiveFunc})
}

var EmptyFunc func() = func() {}

func NewQueueListener(queueListener *QueueListener) ReceiveQueueListener {

	if queueListener.Init == nil {
		queueListener.Init = EmptyFunc
	}
	if queueListener.Receive == nil {
		queueListener.Receive = func(item interface{}) {}
	}
	if queueListener.Empty == nil {
		queueListener.Empty = EmptyFunc
	}
	if queueListener.Shutdown == nil {
		queueListener.Shutdown = EmptyFunc
	}
	if queueListener.Idle == nil {
		queueListener.Idle = EmptyFunc
	}
	if queueListener.StartBatch == nil {
		queueListener.StartBatch = EmptyFunc
	}
	if queueListener.Limit == nil {
		queueListener.Limit = EmptyFunc
	}
	return &BaseReceiveQueueListener{
		init:queueListener.Init,
		receive:queueListener.Receive,
		empty:queueListener.Empty,
		shutdown:queueListener.Shutdown,
		idle:queueListener.Idle,
		startBatch:queueListener.StartBatch,
		limit:queueListener.Limit,
	}
}

type QueueListener struct {
	Init       func()
	Receive    func(item interface{})
	Empty      func()
	Shutdown   func()
	Idle       func()
	StartBatch func()
	Limit      func()
>>>>>>> master
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

func (l *BaseReceiveQueueListener) Limit() {
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
<<<<<<< HEAD
func (*BaseReceiveQueueListener) StartBatch() {
}
=======
func (l *BaseReceiveQueueListener) StartBatch() {
	l.startBatch()
}
>>>>>>> master
