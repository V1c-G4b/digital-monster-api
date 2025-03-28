package entity

type EventDispatcher struct {
	handlers map[string][]func(event DomainEvent)
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]func(event DomainEvent)),
	}
}

func (d *EventDispatcher) Register(eventName string, handler func(event DomainEvent)) {
	d.handlers[eventName] = append(d.handlers[eventName], handler)
}

func (d *EventDispatcher) Dispatch(event DomainEvent) {
	if handlers, ok := d.handlers[event.EventName()]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
