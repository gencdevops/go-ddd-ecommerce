package domain

import "time"

type Event interface {
	Name() string
	Timestamp() time.Time
}

type BaseEvent struct {
	eventName string
	time      time.Time
}

func NewBaseEvent(name string) BaseEvent {
	return BaseEvent{
		eventName: name,
		time:      time.Now(),
	}
}

func (e BaseEvent) Name() string {
	return e.eventName
}

func (e BaseEvent) Timestamp() time.Time {
	return e.time
}
