package webhook

type IAMObjectType string

const (
	IAMObjectTypeTypeWorkspace IAMObjectType = "account.Workspace"
	IAMObjectTypeUser          IAMObjectType = "account.User"
	IAMObjectTypeUserWorkspace IAMObjectType = "account.UserWorkspace"
	// TODO complete this list "later"
)

type EventAction string

const (
	EventActionCreated EventAction = "created"
	EventActionUpdated EventAction = "updated"
	EventActionDeleted EventAction = "deleted"
)

type Event struct {
	Object     IAMObject
	Action     EventAction
	ObjectType IAMObjectType
}
