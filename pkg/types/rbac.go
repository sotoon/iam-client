package types

type RoleBindingReq struct {
	Items map[string]string `json:"items,omitempty"`
}
type RuleReq struct {
	Name    string   `json:"name" validate:"required"`
	Actions []string `json:"actions" validate:"required,gte=1"`
	Object  string   `json:"object" validate:"required"`
	Deny    bool     `json:"deny"`
}
