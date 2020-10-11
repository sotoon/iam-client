package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// Client represents bepa client interface
type Client interface {
	GetWorkspaces() ([]*types.Workspace, error)
	GetWorkspaceByName(name string) (*types.Workspace, error)
	GetWorkspace(uuid *uuid.UUID) (*types.Workspace, error)
	CreateWorkspace(name string) (*types.Workspace, error)
	DeleteWorkspace(uuid *uuid.UUID) error
	GetWorkspaceUsers(uuid *uuid.UUID) ([]*types.User, error)
	GetWorkspaceRoles(uuid *uuid.UUID) ([]*types.Role, error)
	GetWorkspaceRules(uuid *uuid.UUID) ([]*types.Rule, error)
	AddUserToWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	RemoveUserFromWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	SetConfigDefaultWorkspace(uuid *uuid.UUID) error

	CreateRole(roleName string, workspaceUUID *uuid.UUID) (*types.Role, error)
	GetRole(roleUUID, workspaceUUID *uuid.UUID) (*types.Role, error)
	GetRoleByName(roleName, workspaceName string) (*types.Role, error)
	GetAllRoles() ([]*types.Role, error)
	GetRoleUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.User, error)
	GetRoleRules(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Rule, error)
	DeleteRole(roleUUID, workspaceUUID *uuid.UUID) error
	BindRoleToUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error

	GetRule(ruleUUID, workspaceUUID *uuid.UUID) (*types.Rule, error)
	GetRuleByName(ruleName, workspaceName string) (*types.Rule, error)
	CreateRule(ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*types.Rule, error)
	GetAllRules() ([]*types.Rule, error)
	GetAllUserRules(userUUID *uuid.UUID) ([]*types.Rule, error)
	BindRuleToRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	UnbindRuleFromRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	GetRuleRoles(ruleUUID, workspaceUUID *uuid.UUID) ([]*types.Role, error)

	CreateUser(userName, email, password string) (*types.User, error)
	GetUser(userUUID *uuid.UUID) (*types.User, error)
	GetMySelf() (*types.User, error)
	DeleteMySelf() error
	GetUserByName(userName string, workspaceUUID *uuid.UUID) (*types.User, error)
	GetUsers() ([]*types.User, error)
	DeleteUser(userUUID *uuid.UUID) error
	UpdateUser(userUUID *uuid.UUID, name, email, password string) error
	SetMyPassword(password string) error
	SetMyEmail(email string) error
	SetMyName(name string) error
	GetSecret(userUUID *uuid.UUID) (*types.UserSecret, error)
	RevokeSecret(userUUID *uuid.UUID) error
	InviteUser(workspaceUUID *uuid.UUID, email string) (*types.InvitationInfo, error)
	JoinByInvitationToken(server, name, password, invitationToken string) (*types.User, error)
	GetMyWorkspaces() ([]*types.Workspace, error)
	GetUserRoles(userUUID *uuid.UUID) ([]*types.RoleBinding, error)
	CreateUserTokenByCreds(email, password string) (*types.UserToken, error)
	SetConfigDefaultUserData(context, token, userUUID, email string) error
	SetCurrentContext(context string) error
	SuspendUser(userUUID *uuid.UUID) error
	ActivateUser(userUUID *uuid.UUID) error

	CreatePublicKeyForDefaultUser(title, keyType, key string) (*PublicKey, error)
	GetOneDefaultUserPublicKey(publicKeyUUID *uuid.UUID) (*PublicKey, error)
	GetAllDefaultUserPublicKeys() ([]*PublicKey, error)
	DeleteDefaultUserPublicKey(publicKeyUUID *uuid.UUID) error
	CreatePublicKeyFromFileForDefaultUser(title, fileAdd string) (*PublicKey, error)
	VerifyPublicKey(keyType string, key string, workspace_uuid string, username string, hostname string) (bool, error)

	Authorize(identity, action, object string) error
	Identify(token string) (*types.UserRes, error)

	Do(method, path string, req interface{}, resp interface{}) error
	SetAccessToken(token string)
	SetUser(userUUID string)

	CreateTokenWithToken(secret string) (*types.UserToken, error)
	GetUserToken(user_token_uuid string) (*types.UserToken, error)
	GetAllUserToken() (*[]types.UserToken, error)
	DeleteUserToken(user_token_uuid string) error

	GetAllServices() (*[]types.Service, error)
	GetService(name string) (*types.Service, error)

	DeleteServiceUserToken(serviceUserUUID, workspaceUUID, serviceUserTokenUUID *uuid.UUID) error
	GetAllServiceUserToken(serviceUserUUID, workspaceUUID *uuid.UUID) (*[]types.ServiceUserToken, error)
	CreateServiceUserToken(serviceUserUUID, workspaceUUID *uuid.UUID) (*types.ServiceUserToken, error)
	CreateServiceUser(serviceUserName string, workspace *uuid.UUID) (*types.ServiceUser, error)
	GetServiceUserByName(workspaceName string, serviceUserName string, userUUID *uuid.UUID) (*types.ServiceUser, error)
	DeleteServiceUser(workspaceUUID, serviceUserUUID *uuid.UUID) error
	GetServiceUsers(workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error)
	GetServiceUser(workspaceUUID,serviceUserUUID *uuid.UUID) (*types.ServiceUser, error)
	BindRoleToServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error
	GetRoleServiceUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error)

	GetGroup(workspaceUUID, groupUUID *uuid.UUID) (*types.Group, error)
	GetAllGroups(workspaceUUID *uuid.UUID) ([]*types.Group, error)
	DeleteGroup(workspaceUUID, groupUUID *uuid.UUID) error
	GetGroupByName(workspaceName string, groupName string, userUUID *uuid.UUID) (*types.Group, error)
	CreateGroup(groupName string, workspace *uuid.UUID) (*types.Group, error)
	GetGroupUser(workspaceUUID, groupUUID, userUUID *uuid.UUID) (*types.User, error)
	GetAllGroupUsers(workspaceUUID, groupUUID *uuid.UUID) ([]*types.User, error)
	UnbindUserFromGroup(workspaceUUID, groupUUID, userUUID *uuid.UUID) error
	BindGroup(groupName string, workspace, groupUUID, userUUID *uuid.UUID) (*types.GroupUserRes, error)
	GetRoleGroups(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Group, error)
	BindRoleToGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error
}
