package webhook

type EventHandler func(Event) error

type EventListener struct {
	ObjectType IAMObjectType
	Action     EventAction
	Handle     EventHandler
}

func (el *EventListener) doesMatch(event *Event) bool {
	return event.Action == el.Action && event.ObjectType == el.ObjectType
}
