package event

// Listener 监听器接口
type Listener interface {
	Events() []Event
	Process(Event)
}
