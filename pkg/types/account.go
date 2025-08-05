package types

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	UserTypeUser        = "user"
	UserTypeServiceUser = "service-user"
)

type ThirdParty struct {
	UUID string `json:"uuid"`
	Tag  string `json:"tag"`
}

type UserRes struct {
	UUID            string              `json:"uuid"`
	Name            string              `json:"name"`
	Email           string              `json:"email"`
	UserType        string              `json:"user_type"`
	IsSuspended     bool                `json:"is_suspended"`
	Workspaces      []string            `json:"workspaces"`
	CreatedAt       time.Time           `json:"created_at,omitempty"`
	UpdatedAt       time.Time           `json:"updated_at,omitempty"`
	InvitationToken string              `json:"invitation_token,omitempty"`
	Items           []map[string]string `json:"items,omitempty"`
	ThirdParty      ThirdParty          `json:"third_party,omitempty"`
}

type UserTokenReq struct {
	Name      string     `json:"name"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Secret    string     `json:"secret" validate:"required"`
	UserType  string     `json:"user_type"`
}
type IdentifyAndAuthorizeReq struct {
	Token  string `json:"user_token"`
	Object string `json:"object"`
	Action string `json:"action"`
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

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}

type ChangePasswordReq struct {
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
	UUID         *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name         string     `json:"name"`
	Namespace    string     `json:"namespace"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Organization *uuid.UUID `json:"organization" faker:"uuidObject"`
	IsSuspended  bool       `json:"is_suspended"`
}

type WorkspaceWithOrganization struct {
	UUID         *uuid.UUID    `json:"uuid" faker:"uuidObject"`
	Name         string        `json:"name"`
	Namespace    string        `json:"namespace"`
	Organization *Organization `json:"organization"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type Organization struct {
	UUID           *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name           string     `json:"name_en"`
	EnterpriseName string     `json:"enterprise_name"`
	EconomicCode   string     `json:"economic_code"`
	NationalId     string     `json:"national_id"`
}

type AuthnChallengeRequiredResponse struct {
	ChallengeToken string `json:"challenge_token"`
	ChallengeType  string `json:"challenge_type"`
}

type AuthnChallengeRequest struct {
	ChallengeToken  string `json:"challenge_token"`
	ChallengeAnswer string `json:"challenge_answer"`
}

func (r *AuthnChallengeRequiredResponse) Error() string {
	return fmt.Sprintf("challenge of type '%s' required", r.ChallengeType)
}

type UserToken struct {
	UUID         string     `json:"uuid"`
	Name         string     `json:"name"`
	User         string     `json:"user"`
	Secret       string     `json:"secret"`
	IsHashed     bool       `json:"is_hashed"`
	Active       bool       `json:"active"`
	LastAccessAt *time.Time `json:"last_access_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
}

type User struct {
	UUID                *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name                string     `json:"name"`
	Email               string     `json:"email"`
	InvitationToken     string     `json:"invitation_token,omitempty"`
	Groups              []Group    `json:"groups"`
	Roles               []Role     `json:"roles"`
	EmailVerified       bool       `json:"email_verified"`
	PhoneNumber         string     `json:"phone_number,omitempty"`
	PhoneNumberVerified bool       `json:"phone_number_verified"`
	FirstName           string     `json:"first_name,omitempty"`
	LastName            string     `json:"last_name,omitempty"`
	Birthday            string     `json:"birthday,omitempty"`
	IsSuspended         bool       `json:"is_suspended"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	UserType            string     `json:"user_type"`
	IsOtpEnabled        bool       `json:"is_otp_enabled"`
}

type UserWithCompactRole struct {
	UUID                *uuid.UUID             `json:"uuid" faker:"uuidObject"`
	Name                string                 `json:"name"`
	Email               string                 `json:"email"`
	InvitationToken     string                 `json:"invitation_token,omitempty"`
	Groups              []Group                `json:"groups"`
	Roles               []RoleCompactWorkspace `json:"roles"`
	EmailVerified       bool                   `json:"email_verified"`
	PhoneNumber         string                 `json:"phone_number,omitempty"`
	PhoneNumberVerified bool                   `json:"phone_number_verified"`
	FirstName           string                 `json:"first_name,omitempty"`
	LastName            string                 `json:"last_name,omitempty"`
	Birthday            string                 `json:"birthday,omitempty"`
	IsSuspended         bool                   `json:"is_suspended"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	UserType            string                 `json:"user_type"`
	IsOtpEnabled        bool                   `json:"is_otp_enabled"`
}
type Group struct {
	UUID               *uuid.UUID             `json:"uuid" faker:"uuidObject"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	Workspace          Workspace              `json:"workspace"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	UsersNumber        int                    `json:"users_number,omitempty"`
	ServiceUsersNumber int                    `json:"service_users_number,omitempty"`
	Roles              []RoleCompactWorkspace `json:"roles,omitempty"`
}

type GroupUser struct {
	User  *User  `json:"user"`
	Group *Group `json:"group"`
}

type GroupServiceUser struct {
	Group       *Group       `json:"group"`
	ServiceUser *ServiceUser `json:"service_user"`
}
type GroupReq struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
}

type GroupUpdateReq struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Workspace   WorkspaceUpdateReq `json:"workspace"`
}

type WorkspaceUpdateReq struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Organization string `json:"organization"`
	IsSuspended  bool   `json:"is_suspended"`
}
type GroupRes struct {
	UUID          *uuid.UUID `json:"uuid" faker:"uuidObject"`
	Name          string     `json:"name"`
	WorkspaceUUID string     `json:"workspace"`
	Descriotion   string     `json:"description"`
}
type GroupUserRes struct {
	Group string `json:"group"`
	User  string `json:"user"`
}

type AutoGenerated []struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ServiceUser struct {
	UUID        *uuid.UUID  `json:"uuid" faker:"uuidObject"`
	Groups      []Group     `json:"groups"`
	Description string      `json:"description"`
	ThirdParty  *ThirdParty `json:"third_party,omitempty"`
	Name        string      `json:"name"`
	Workspace   *uuid.UUID  `json:"workspace"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
	Roles       []*Role     `json:"roles,omitempty"`
}

type ServiceUserWithCompactRole struct {
	UUID        *uuid.UUID              `json:"uuid" faker:"uuidObject"`
	Groups      []Group                 `json:"groups"`
	Description string                  `json:"description"`
	ThirdParty  *ThirdParty             `json:"third_party,omitempty"`
	Name        string                  `json:"name"`
	Workspace   *uuid.UUID              `json:"workspace"`
	CreatedAt   string                  `json:"created_at,omitempty"`
	UpdatedAt   string                  `json:"updated_at,omitempty"`
	Roles       []*RoleCompactWorkspace `json:"roles,omitempty"`
}
type ServiceUserReq struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
}

type ServiceUserToken struct {
	UUID        *uuid.UUID `json:"uuid" faker:"uuidObject"`
	ServiceUser string     `json:"service_user"`
	Secret      string     `json:"secret"`
	Name        string     `json:"name"`
	IsHashed    bool       `json:"is_hashed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

type ServiceUserPublicKey struct {
	UUID        *uuid.UUID `json:"uuid" faker:"uuidObject"`
	ServiceUser string     `json:"service_user"`
	PublicKey   string     `json:"public_key"`
	Name        string     `json:"name"`
	CreatedAt   string     `json:"created_at,omitempty"`
	Title       string     `json:"title"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Key         string     `json:"key"`
	Type        string     `json:"type"`
}

type InvitationInfo struct {
	Token string `json:"invitation_token"`
}

type UserOtp struct {
	UUID      *uuid.UUID `json:"uuid" faker:"uuidObject"`
	User      string     `json:"user"`
	Secret    string     `json:"secret,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
	UpdatedAt string     `json:"updated_at,omitempty"`
}

type UserSecret struct {
	Secret string `json:"secret"`
}
type Service struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type PublicKey struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Key       string    `json:"key"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt any       `json:"deleted_at"`
	Type      string    `json:"type"`
	PublicKey string    `json:"public_key"`
}

type KiseSecret struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Secret      string `json:"secret"`
	User        string `json:"user"`
	ServiceUser string `json:"service_user,omitempty"`
}

type OpenIDToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type ThirdPartyBulkRefreshToken struct {
	UUID            string     `json:"uuid"`
	RefreshToken    string     `json:"refresh_token"`
	ExpiresAt       *time.Time `json:"expires_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	ServiceUserUUID string     `json:"service_user_uuid"`
	ThirdPartyUUID  string     `json:"third_party_uuid"`
	WorkspaceUUID   string     `json:"workspace_uuid"`
}

type ThirdPartyAccessToken struct {
	UUID             string     `json:"uuid"`
	AccessToken      string     `json:"access_token"`
	ExpiresAt        *time.Time `json:"expires_at,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	ThirdPartyUUID   string     `json:"third_party_uuid"`
	OrganizationUUID string     `json:"organization_uuid"`
}
type BackupKey struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	Key       string `json:"key"`
	Type      string `json:"type"`
	Workspace string `json:"workspace"`
}
type BackupKeyReq struct {
	Title string `json:"title"`
	Key   string `json:"key" validate:"required"`
}

type WebhookWorkspaceOrganization struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type WebhookWorkspace struct {
	UUID         uuid.UUID    `json:"uuid"`
	Name         string       `json:"name"`
	IsSuspended  bool         `json:"is_suspended"`
	Organization Organization `json:"organization"`

	// TODO fix parsing error of CreatedAt and UpdatedAt
	// CreatedAt    time.Time    `json:"created_at"`
	// UpdatedAt    time.Time    `json:"updated_at"`
}

type WebhookUser struct {
	UUID                  uuid.UUID `json:"uuid"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	IsEmailVerified       bool      `json:"email_verified"`
	PhoneNumber           string    `json:"phone_number"`
	IsPhoneNumberVerified bool      `json:"phone_number_verified"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	Birthday              string    `json:"birthday"`
	IsSuspended           bool      `json:"is_suspended"`

	// TODO fix parsing error of CreatedAt and UpdatedAt
	// CreatedAt             time.Time `json:"created_at"`
	// UpdatedAt             time.Time `json:"updated_at"`
}

type WebhookUserWorkspaceRelation struct {
	UserUUID      uuid.UUID `json:"user_uuid"`
	WorkspaceUUID uuid.UUID `json:"workspace_uuid"`
	IsSuspended   bool      `json:"is_suspended"`
}
