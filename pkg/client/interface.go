package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// Client represents bepa client interface
type Client interface {
	GetWorkspaces() ([]*Workspace, error)
	GetWorkspaceByName(name string) (*Workspace, error)
	GetWorkspace(uuid *uuid.UUID) (*Workspace, error)
	CreateWorkspace(name string) (*Workspace, error)
	DeleteWorkspace(uuid *uuid.UUID) error
	GetWorkspaceUsers(uuid *uuid.UUID) ([]*User, error)
	GetWorkspaceRoles(uuid *uuid.UUID) ([]*Role, error)
	GetWorkspaceRules(uuid *uuid.UUID) ([]*Rule, error)
	AddUserToWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	RemoveUserFromWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	SetConfigDefaultWorkspace(uuid *uuid.UUID) error

	CreateRole(roleName string, workspaceUUID *uuid.UUID) (*Role, error)
	GetRole(roleUUID, workspaceUUID *uuid.UUID) (*Role, error)
	GetRoleByName(roleName, workspaceName string) (*Role, error)
	GetAllRoles() ([]*Role, error)
	GetRoleUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*User, error)
	GetRoleRules(roleUUID, workspaceUUID *uuid.UUID) ([]*Rule, error)
	DeleteRole(roleUUID, workspaceUUID *uuid.UUID) error
	BindRoleToUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromUser(workspaceUUID, roleUUID, userUUID *uuid.UUID) error

	GetRule(ruleUUID, workspaceUUID *uuid.UUID) (*Rule, error)
	GetRuleByName(ruleName, workspaceName string) (*Rule, error)
	CreateRule(ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*Rule, error)
	GetAllRules() ([]*Rule, error)
	GetAllUserRules(userUUID *uuid.UUID) ([]*Rule, error)
	BindRuleToRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	UnbindRuleFromRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	GetRuleRoles(ruleUUID, workspaceUUID *uuid.UUID) ([]*Role, error)

	CreateUser(userName, email, password string) (*User, error)
	GetUser(userUUID *uuid.UUID) (*User, error)
	GetMySelf() (*User, error)
	DeleteMySelf() error
	GetUserByName(userName string, workspaceUUID *uuid.UUID) (*User, error)
	GetUsers() ([]*User, error)
	DeleteUser(userUUID *uuid.UUID) error
	UpdateUser(userUUID *uuid.UUID, name, email, password string) error
	SetMyPassword(password string) error
	SetMyEmail(email string) error
	SetMyName(name string) error
	GetSecret(userUUID *uuid.UUID) (*UserSecret, error)
	RevokeSecret(userUUID *uuid.UUID) error
	InviteUser(workspaceUUID *uuid.UUID, email string) (*InvitationInfo, error)
	JoinByInvitationToken(server, name, password, invitationToken string) (*User, error)
	GetMyWorkspaces() ([]*Workspace, error)
	GetUserRoles(userUUID *uuid.UUID) ([]*RoleBinding, error)
	CreateUserTokenByCreds(email, password string) (*UserToken, error)
	SetConfigDefaultUserData(context, token, userUUID, email string) error
	SetCurrentContext(context string) error

	CreatePublicKeyForDefaultUser(title, keyType, key string) (*PublicKey, error)
	GetOneDefaultUserPublicKey(publicKeyUUID *uuid.UUID) (*PublicKey, error)
	GetAllDefaultUserPublicKeys() ([]*PublicKey, error)
	DeleteDefaultUserPublicKey(publicKeyUUID *uuid.UUID) error
	CreatePublicKeyFromFileForDefaultUser(title, fileAdd string) (*PublicKey, error)

	Authorize(identity, action, object string) error
	Identify(token string) (*types.UserRes, error)

	Do(method, path string, req interface{}, resp interface{}) error
	SetAccessToken(token string)
	SetUser(userUUID string)
	SetServerURL(serverURL string) error
	GetServerURL() string
}
