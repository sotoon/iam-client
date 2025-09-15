package client

import (
	"log"
	"net/http"
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client/interceptor"
	"github.com/sotoon/iam-client/pkg/types"
)

// Client represents iam client interface
type Client interface {
	AddInterceptor(i interceptor.ClientInterceptor)
	ProcessRequest(httpRequest *http.Request, successCode int, id string) (*http.Response, error)
	// IsHealthy reports if the client can connect to the IAM Server
	// For ReliableClient returns (true, nil) if it can connect to `at least one` healthy IAM Server endpoint
	// For SimpleClient returns (true, nil) if it can connect to `the exactly one` IAM Server Endpoint and it is healthy
	IsHealthy() (bool, error)

	GetOrganizations() ([]*types.Organization, error)
	GetOrganization(*uuid.UUID) (*types.Organization, error)
	GetOrganizationWorkspaces(*uuid.UUID) ([]*types.Workspace, error)
	GetOrganizationWorkspace(*uuid.UUID, *uuid.UUID) (*types.Workspace, error)

	GetWorkspaces() ([]*types.Workspace, error)
	GetWorkspaceByName(name string) (*types.Workspace, error)
	GetWorkspaceByNameAndOrgName(name string, organizationName string) (*types.WorkspaceWithOrganization, error)
	GetWorkspace(uuid *uuid.UUID) (*types.Workspace, error)
	CreateWorkspace(name string) (*types.Workspace, error)
	DeleteWorkspace(uuid *uuid.UUID) error
	GetWorkspaceUsers(uuid *uuid.UUID) ([]*types.User, error)
	GetWorkspaceRoles(uuid *uuid.UUID) ([]*types.Role, error)
	GetWorkspaceRules(uuid *uuid.UUID) ([]*types.Rule, error)
	AddUserToWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	RemoveUserFromWorkspace(userUUID, workspaceUUID *uuid.UUID) error
	SetConfigDefaultWorkspace(uuid *uuid.UUID) error
	GetWorkspaceServices(workspaceUUID uuid.UUID) ([]types.Service, error)

	CreateRole(roleName, description string, workspaceUUID *uuid.UUID) (*types.RoleWithCompactWorkspace, error)
	UpdateRole(roleUUID *uuid.UUID, roleName string, workspaceUUID *uuid.UUID) (*types.Role, error)
	GetRole(roleUUID, workspaceUUID *uuid.UUID) (*types.RoleRes, error)
	GetRoleByName(roleName, workspaceName string) (*types.RoleRes, error)
	GetAllRoles() ([]*types.Role, error)
	GetRoleUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.User, error)
	GetRoleRules(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Rule, error)
	DeleteRole(roleUUID, workspaceUUID *uuid.UUID) error
	BindRoleToUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error
	GetBindedRoleToUserItems(workspaceUUID, roleUUID, userUUID *uuid.UUID) (map[string]string, error)
	GetBindedRoleToServiceUserItems(workspaceUUID, roleUUID, userUUID *uuid.UUID) (map[string]string, error)
	GetBindedRoleToGroupItems(workspaceUUID, roleUUID, userUUID *uuid.UUID) (map[string]string, error)

	GetRule(ruleUUID, workspaceUUID *uuid.UUID) (*types.Rule, error)
	GetRuleByName(ruleName, workspaceName string) (*types.Rule, error)
	CreateRule(ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*types.Rule, error)
	DeleteRule(ruleUUID, workspaceUUID *uuid.UUID) error
	GetAllRules() ([]*types.Rule, error)
	GetAllUserRules(userUUID *uuid.UUID) ([]*types.Rule, error)
	BindRuleToRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	UnbindRuleFromRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error
	GetRuleRoles(ruleUUID, workspaceUUID *uuid.UUID) ([]*types.Role, error)
	UpdateRule(ruleUUID *uuid.UUID, ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*types.Rule, error)

	CreateUser(userName, email, password string) (*types.User, error)
	GetUser(userUUID *uuid.UUID) (*types.User, error)
	GetMySelf() (*types.User, error)
	DeleteMySelf() error
	GetUserByEmail(email string, workspaceUUID *uuid.UUID) (*types.User, error)
	GetUserByName(userName string, workspaceUUID *uuid.UUID) (*types.User, error)
	GetUsers() ([]*types.User, error)
	DeleteUser(userUUID *uuid.UUID) error
	GetWorkspaceUserList(workspaceUUID uuid.UUID) ([]*types.UserWithCompactRole, error)
	GetWorkspaceUserDetail(workspaceUUID, userUUID uuid.UUID) (*types.UserWithCompactRole, error)
	GetUserOtp(userUUID uuid.UUID) (*types.UserOtp, error)
	CreateUserOtp(userUUID uuid.UUID) (*types.UserOtp, error)
	DeleteUserOtp(userUUID uuid.UUID) error
	UpdateUser(userUUID *uuid.UUID, name, email, password string) error
	SetMyPassword(password string) error
	SetMyEmail(email string) error
	SetMyName(name string) error
	GetSecret(userUUID *uuid.UUID) (*types.UserSecret, error)
	RevokeSecret(userUUID *uuid.UUID) error
	SuspendUserInWorkspace(workspaceUUID *uuid.UUID, userUUID *uuid.UUID) error
	ActivateUserInWorkspace(workspaceUUID *uuid.UUID, userUUID *uuid.UUID) error
	InviteUser(workspaceUUID *uuid.UUID, email string) (*types.InvitationInfo, error)
	JoinByInvitationToken(name, password, invitationToken string) (*types.User, error)
	ResetPassword(email string) error
	ChangePassword(token, password string) error
	GetMyWorkspaces() ([]*types.WorkspaceWithOrganization, error)
	GetUserRoles(userUUID *uuid.UUID) ([]*types.RoleBinding, error)
	CreateMyUserTokenWithTokenByCreds(email, password string) (*types.UserToken, error)
	SetConfigDefaultUserData(context, token, userUUID, email string) error
	SetCurrentContext(context string) error
	SuspendUser(userUUID *uuid.UUID) error
	ActivateUser(userUUID *uuid.UUID) error

	CreateMyUserPublicKey(title, keyType, key string) (*types.PublicKey, error)
	GetOneDefaultUserPublicKey(publicKeyUUID *uuid.UUID) (*types.PublicKey, error)
	GetAllMyUserPublicKeyList() ([]*types.PublicKey, error)
	DeleteMyUserPublicKey(publicKeyUUID *uuid.UUID) error
	CreatePublicKeyFromFileForDefaultUser(title, fileAdd string) (*types.PublicKey, error)
	VerifyPublicKey(keyType string, key string, workspaceUUID string, username string, hostname string) (bool, error)

	GetUserKiseSecrets(userUUID *uuid.UUID, workspaceUUID *uuid.UUID) ([]*types.KiseSecret, error)
	CreateUserKiseSecret(userUUID *uuid.UUID, workspaceUUID *uuid.UUID, title string) (*types.KiseSecret, error)
	DeleteUserKiseSecret(KiseSecretUUID *uuid.UUID) error
	CreateKiseSecretForDefaultUser() (*types.KiseSecret, error)
	GetServiceUserKiseSecrets(workspaceUUID uuid.UUID) ([]*types.KiseSecret, error)
	CreateServiceUserKiseSecret(workspaceUUID, serviceUserUUID uuid.UUID, title string) (*types.KiseSecret, error)
	DeleteServiceUserKiseSecret(workspaceUUID, serviceUserUUID, kiseSecretUUID uuid.UUID) error

	GetThirdPartyBulkRefreshTokens(workspaceUUID, thirdPartyUUID, serviceUserUUID uuid.UUID) ([]*types.ThirdPartyBulkRefreshToken, error)
	CreateThirdPartyBulkRefreshToken(workspaceUUID, thirdPartyUUID, serviceUserUUID uuid.UUID, refreshToken string, expiresAt *time.Time) (*types.ThirdPartyBulkRefreshToken, error)

	GetThirdPartyAccessTokens(organizationUUID, thirdPartyUUID uuid.UUID) ([]*types.ThirdPartyAccessToken, error)
	CreateThirdPartyAccessToken(organizationUUID, thirdPartyUUID uuid.UUID, accessToken string, expiresAt *time.Time) (*types.ThirdPartyAccessToken, error)

	Authorize(identity, userType, action, rriObject string) error
	Identify(token string) (*types.UserRes, error)
	IdentifyAndAuthorize(token, action, rriObject string) error

	Do(method, path string, successCode int, req interface{}, resp interface{}) error
	SetLogger(logger *log.Logger)
	SetAccessToken(token string)
	SetDefaultWorkspace(workspace string)
	SetUser(userUUID string)

	CreateMyUserToken(name string, expiresAt *time.Time) (*types.UserToken, error)
	GetMyUserToken(UserTokenUUID *uuid.UUID) (*types.UserToken, error)
	GetAllMyUserTokenList() (*[]types.UserToken, error)
	DeleteMyUserToken(UserTokenUUID *uuid.UUID) error

	GetService(name string) (*types.Service, error)

	BindUserToGroup(workspaceUUID, groupUUID, userUUID *uuid.UUID) error
	UpdateServiceUser(workspaceUUID, serviceUserUUID uuid.UUID, name, description string) (*types.ServiceUser, error)
	DeleteServiceUserToken(serviceUserUUID, workspaceUUID, serviceUserTokenUUID *uuid.UUID) error
	GetWorkspaceServiceUserTokenList(serviceUserUUID, workspaceUUID *uuid.UUID) (*[]types.ServiceUserToken, error)
	CreateServiceUserToken(serviceUserUUID, workspaceUUID *uuid.UUID, name string, expiresAt *time.Time) (*types.ServiceUserToken, error)
	CreateServiceUser(serviceUserName, description string, workspace *uuid.UUID) (*types.ServiceUser, error)
	GetServiceUserByName(workspaceName string, serviceUserName string) (*types.ServiceUser, error)
	DeleteServiceUser(workspaceUUID, serviceUserUUID *uuid.UUID) error
	GetServiceUsers(workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error)
	GetServiceUser(workspaceUUID, serviceUserUUID *uuid.UUID) (*types.ServiceUser, error)
	GetWorkspaceServiceUserList(workspaceUUID uuid.UUID) ([]*types.ServiceUserWithCompactRole, error)
	GetWorkspaceServiceUserDetail(workspaceUUID, serviceUserUUID uuid.UUID) (*types.ServiceUserWithCompactRole, error)
	GetWorkspaceServiceUserPublicKeyList(workspaceUUID, serviceUserUUID uuid.UUID) ([]*types.ServiceUserPublicKey, error)
	CreateServiceUserPublicKey(workspaceUUID, serviceUserUUID uuid.UUID, name, publicKey string) (*types.ServiceUserPublicKey, error)
	DeleteServiceUserPublicKey(workspaceUUID, serviceUserUUID, publicKeyUUID uuid.UUID) error
	BindRoleToServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error
	GetRoleServiceUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error)
	BulkAddServiceUsersToRole(workspaceUUID, roleUUID uuid.UUID, serviceUserUUIDs []uuid.UUID) error
	BulkAddUsersToRole(workspaceUUID, roleUUID uuid.UUID, userUUIDs []uuid.UUID) error
	BulkAddRulesToRole(workspaceUUID, roleUUID uuid.UUID, ruleUUIDs []uuid.UUID) error

	GetGroup(workspaceUUID, groupUUID *uuid.UUID) (*types.Group, error)
	GetAllGroups(workspaceUUID *uuid.UUID) ([]*types.Group, error)
	GetWorkspaceGroupList(workspaceUUID uuid.UUID) ([]*types.Group, error)
	GetWorkspaceGroupDetail(workspaceUUID, groupUUID uuid.UUID) (*types.Group, error)
	GetWorkspaceGroupRoleList(workspaceUUID, groupUUID uuid.UUID) ([]*types.Role, error)
	BulkAddUsersToGroup(workspaceUUID, groupUUID uuid.UUID, userUUIDs []uuid.UUID) ([]*types.GroupUser, error)
	BulkAddServiceUsersToGroup(workspaceUUID, groupUUID uuid.UUID, serviceUserUUIDs []uuid.UUID) ([]*types.GroupServiceUser, error)
	BulkAddRolesToGroup(workspaceUUID, groupUUID uuid.UUID, rolesWithItems []types.RoleWithItems) error
	DeleteGroup(workspaceUUID, groupUUID *uuid.UUID) error
	GetGroupByName(workspaceName string, groupName string) (*types.Group, error)
	CreateGroup(groupName, description string, workspace *uuid.UUID) (*types.GroupRes, error)
	UpdateGroup(workspaceUUID, groupUUID uuid.UUID, name, description *string, workspaceInfo *types.WorkspaceUpdateReq) error
	GetGroupUser(workspaceUUID, groupUUID, userUUID *uuid.UUID) (*types.User, error)
	GetAllGroupUserList(workspaceUUID, groupUUID *uuid.UUID) ([]*types.User, error)
	GetAllGroupServiceUserList(workspaceUUID, groupUUID *uuid.UUID) ([]*types.ServiceUser, error)
	UnbindUserFromGroup(workspaceUUID, groupUUID, userUUID *uuid.UUID) error
	BindGroup(groupName string, workspace, groupUUID, userUUID *uuid.UUID) error
	GetRoleGroups(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Group, error)
	BindRoleToGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error
	UnbindRoleFromGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error
	BindServiceUserToGroup(worspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) error
	UnbindServiceUserFromGroup(worspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) error
	GetGroupServiceUser(worspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) (*types.ServiceUser, error)

	GetServerURL() string

	GetAllDefaultBackupKeys() ([]*types.BackupKey, error)
	GetOneDefaultBackupKey(BackupKeyUUID *uuid.UUID) (*types.BackupKey, error)
	DeleteDefaultWorkspaceBackupKey(backupKeyUUID *uuid.UUID) error
	CreateBackupKeyForDefaultWorkspace(title, keyType, key string) (*types.BackupKey, error)
	CreateBackupKeyFromFileForDefaultUser(title, fileAdd string) (*types.BackupKey, error)

	GetBaseURL() (*url.URL, error)
}

type Cache interface {
	Get(string) (interface{}, bool)
	Set(string, interface{}, time.Duration)
}
