package types

import "time"

type UserRes struct {
	UUID            string              `json:"uuid"`
	Name            string              `json:"name"`
	Email           string              `json:"email"`
	IsSuspended     bool                `json:"is_suspended"`
	CreatedAt       time.Time           `json:"created_at,omitempty"`
	UpdatedAt       time.Time           `json:"updated_at,omitempty"`
	InvitationToken string              `json:"invitation_token,omitempty"`
	Items           []map[string]string `json:"items,omitempty"`
}
type UserTokenReq struct {
	Secret string `json:"secret" validate:"required"`
}
type UserUpdateReq struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
}

type UserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserAcceptInvitationReq struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserCanReq struct {
	Path string `json:"path" validate:"required"`
}

type UserSecretRes struct {
	Secret string `json:"secret"`
}

type UserTokenByCredsReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type InviteUserReq struct {
	Email string `json:"email" validate:"required,email"`
}

type RoleReq struct {
	Name      string `json:"name" validate:"required"`
	Workspace string `json:"workspace" validate:"required"`
}

type WorkspaceReq struct {
	Name string `json:"name" validate:"required,rfc1123_label"`
}

type PublicKeyReq struct {
	Title string `json:"title"`
	Key   string `json:"key" validate:"required"`
}

type VerifRes struct {
	Message string `json:"message"`
}
