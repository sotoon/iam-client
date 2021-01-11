package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

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
	Secret   string `json:"secret" validate:"required"`
	UserType string `json:"user_type"`
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

type PublicKeyVerifyReq struct {
	Key            string
	KeyType        string
	Workspace_uuid string `json:"workspace_uuid"`
	Hostname       string `json:"hostname"`
	Email          string `json:"email"`
}
type Workspace struct {
	UUID      *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
type UserToken struct {
	UUID         string     `json:"uuid"`
	User         string     `json:"user"`
	Secret       string     `json:"secret"`
	Active       bool       `json:"active"`
	LastAccessAt *time.Time `json:"last_access_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type User struct {
	UUID            *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	InvitationToken string     `json:"invitation_token,omitempty"`
}
type Group struct {
	UUID      *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name      string     `json:"name"`
	Workspace Workspace  `json:"workspace"`
}
type GroupReq struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
}
type GroupUserRes struct {
	Group string `json:"group"`
	User  string `json:"user"`
}
type ServiceUser struct {
	UUID      *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name      string     `json:"name"`
	Workspace string     `json:"workspace"`
}
type ServiceUserReq struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
}
type ServiceUserToken struct {
	UUID        *uuid.UUID `json:"uuid" faker:"uuidObject"`
	ServiceUser string     `json:"service_user"`
	Secret      string     `json:"secret"`
}

type InvitationInfo struct {
	Token string `json:"invitation_token"`
}

type UserSecret struct {
	Secret string `json:"secret"`
}
type Service struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type PublicKey struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Key   string `json:"key"`
	User  string `json:"user"`
}
