package qbit


func NewListener () ReceiveQueueListener {
	return  &BaseReceiveQueueListener{}
}


func NewListenerReceive (receiveFunc func(item interface{})) ReceiveQueueListener {

	return &BaseReceiveQueueListener{receive:receiveFunc}
}

type BaseReceiveQueueListener struct {
	init func()
	receive func(item interface{})
	empty func()
	shutdown func()
	idle func()
	startBatch func()
}

func (*BaseReceiveQueueListener) Limit() {
}
func (*BaseReceiveQueueListener) Init() {

}
func (l *BaseReceiveQueueListener) Receive(item interface{}) {
	l.receive(item)
}
func (*BaseReceiveQueueListener) Empty() {
}
func (*BaseReceiveQueueListener) Shutdown() {
}
func (*BaseReceiveQueueListener) Idle() {
}
func (*BaseReceiveQueueListener) StartBatch() {
}