package webhook

type EventHandler func(Event) error

type EventListener struct {
	objectType IAMObjectType
	action     EventAction
	handle     EventHandler
}

func (el *EventListener) doesMatch(event *Event) bool {
	return event.Action == el.action && event.Object == el.objectType
}
