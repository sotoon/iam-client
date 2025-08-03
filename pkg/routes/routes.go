package routes

const (
	// healthz
	RouteHealthCheck = "/healthz/"

	//Helper
	RouteUserBulkCan = "/user/{user_uuid}/bulk-can/workspace/{workspace_uuid}/"

	// Auth
	RouteAuthz                  = "/authz/"
	RouteUserTokenCreateByCreds = "/authn/"
	RouteUserTokenByChallenge   = "/authn/challenge/"
	RouteUserResetPassword      = "/user/reset-password/"
	RouteUserChangePassword     = "/user/change-password/{token}/"
	RouteAcceptInvitation       = "/accept-invitation/{user_invitation_token}/"

	// Groups
	RouteGroupGetOne              = "/workspace/{workspace_uuid}/group/{group_uuid}/"
	RouteGroupDelete              = "/workspace/{workspace_uuid}/group/{group_uuid}/"
	RouteGroupUpdate              = "/workspace/{workspace_uuid}/group/{group_uuid}/"
	RouteGroupGetALL              = "/workspace/{workspace_uuid}/group/"
	RouteGroupCreate              = "/workspace/{workspace_uuid}/group/"
	RouteWorkspaceGroupList       = "/detailed/workspace/{workspace_uuid}/group/"
	RouteWorkspaceGroupDetail     = "/detailed/workspace/{workspace_uuid}/group/{group_uuid}/"
	RouteWorkspaceGroupRoles      = "/workspace/{workspace_uuid}/group/{group_uuid}/role/"
	RouteGroupServiceUserGetALL   = "/workspace/{workspace_uuid}/group/{group_uuid}/service-user/"
	RouteGroupUserGetALL          = "/workspace/{workspace_uuid}/group/{group_uuid}/user/"
	RouteGroupBulkAddUsers        = "/workspace/{workspace_uuid}/group/{group_uuid}/bulk-add-users/"
	RouteGroupBulkAddServiceUsers = "/workspace/{workspace_uuid}/group/{group_uuid}/bulk-add-service-users/"
	RouteGroupBulkAddRoles        = "/workspace/{workspace_uuid}/group/{group_uuid}/bulk-add-roles/"
	RouteGroupUnbindServiceUser   = "/workspace/{workspace_uuid}/group/{group_uuid}/service-user/{service_user_uuid}/"
	RouteGroupBindServiceUser     = "/workspace/{workspace_uuid}/group/{group_uuid}/service-user/{service_user_uuid}/"
	RouteGroupUnbindUser          = "/workspace/{workspace_uuid}/group/{group_uuid}/user/{user_uuid}/"
	RouteGroupBindUser            = "/workspace/{workspace_uuid}/group/{group_uuid}/user/{user_uuid}/"
	RouteGroupGetBindedRole       = "/workspace/{workspace_uuid}/role/{role_uuid}/group/{group_uuid}/"

	//Public Key
	RoutePublicKeyCreate = "/user/{user_uuid}/public-key/"
	RoutePublicKeyGetAll = "/user/{user_uuid}/public-key/"
	RoutePublicKeyDelete = "/user/{user_uuid}/public-key/{public_key_uuid}/"

	//Roles
	RouteRoleCreate              = "/workspace/{workspace_uuid}/role/"
	RouteWorkspaceGetAllRoles    = "/workspace/{workspace_uuid}/role/"
	RouteRoleGetOne              = "/workspace/{workspace_uuid}/role/{role_uuid}/"
	RouteRoleDelete              = "/workspace/{workspace_uuid}/role/{role_uuid}/"
	RouteRoleBulkAddServiceUsers = "/workspace/{workspace_uuid}/role/{role_uuid}/bulk-add-service-users/"
	RouteRoleBulkAddUsers        = "/workspace/{workspace_uuid}/role/{role_uuid}/bulk-add-users/"
	RouteRoleGetAllRules         = "/workspace/{workspace_uuid}/role/{role_uuid}/rule/"
	RouteRoleAppendRule          = "/workspace/{workspace_uuid}/role/{role_uuid}/rule/{rule_uuid}/"
	RouteRoleDropRule            = "/workspace/{workspace_uuid}/role/{role_uuid}/rule/{rule_uuid}/"
	RouteRoleBulkAddRules        = "/workspace/{workspace_uuid}/role/{role_uuid}/bulk-add-rules/"
	RouteRoleGetAllServiceUsers  = "/workspace/{workspace_uuid}/role/{role_uuid}/service-user/"
	RouteServiceUserAppendRole   = "/workspace/{workspace_uuid}/role/{role_uuid}/service-user/{service_user_uuid}/"
	RouteServiceUserDropRole     = "/workspace/{workspace_uuid}/role/{role_uuid}/service-user/{service_user_uuid}/"
	RouteRoleGetAllUsers         = "/workspace/{workspace_uuid}/role/{role_uuid}/user/"
	RouteUserDropRole            = "/workspace/{workspace_uuid}/role/{role_uuid}/user/{user_uuid}/"

	//Rule
	RouteRuleCreate           = "/workspace/{workspace_uuid}/rule/"
	RouteWorkspaceGetAllRules = "/workspace/{workspace_uuid}/rule/"
	RouteRuleUpdate           = "/workspace/{workspace_uuid}/rule/{rule_uuid}/"
	RouteRuleDelete           = "/workspace/{workspace_uuid}/rule/{rule_uuid}/"
	RouteRuleGetOne           = "/workspace/{workspace_uuid}/rule/{rule_uuid}/"
	RouteWorkspaceService     = "/workspace/{workspace_uuid}/service/"
	RouteRuleGetAllRoles      = "/workspace/{workspace_uuid}/rule/{rule_uuid}/role/"

	// Service User
	RouteServiceUserCreate          = "/workspace/{workspace_uuid}/service-user/"
	RouteServiceUserGetALL          = "/workspace/{workspace_uuid}/service-user/"
	RouteServiceUserGetOne          = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/"
	RouteServiceUserDelete          = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/"
	RouteServiceUserDetailList      = "/detailed/workspace/{workspace_uuid}/service-user/"
	RouteServiceUserDetailGetOne    = "/detailed/workspace/{workspace_uuid}/service-user/{service_user_uuid}/"
	RouteServiceUserPublicKeyList   = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/service-user-public-key/"
	RouteServiceUserPublicKeyCreate = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/service-user-public-key/"
	RouteServiceUserPublicKeyDelete = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/service-user-public-key/{public_key_uuid}/"
	RouteServiceUserTokenGetALL     = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/token/"
	RouteServiceUserTokenCreate     = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/token/"
	RouteServiceUserTokenDelete     = "/workspace/{workspace_uuid}/service-user/{service_user_uuid}/token/{service_user_token_uuid}/"

	// SSH Key
	RouteBackupKeyCreate = "/workspace/{workspace_uuid}/backup-key/"
	RouteBackupKeyGetAll = "/workspace/{workspace_uuid}/backup-key/"
	RouteBackupKeyDelete = "/workspace/{workspace_uuid}/backup-key/{backup_key_uuid}/"

	// Token
	RouteUserTokenCreateByToken = "/user/{user_uuid}/user-token/"
	RouteUserTokenGetAll        = "/user/{user_uuid}/user-token/"
	RouteUserTokenDelete        = "/user/{user_uuid}/user-token/{user_token_uuid}/"

	// User
	RouteUserGetAllWorkspaces    = "/user/{user_uuid}/workspace/"
	RouteWorkspaceGetUsers       = "/workspace/{workspace_uuid}/user/"
	RouteUserLeaveWorkspace      = "/workspace/{workspace_uuid}/user/{user_uuid}/"
	RouteUserDetailedList        = "/detailed/workspace/{workspace_uuid}/user/"
	RouteUserDetailedDetail      = "/detailed/workspace/{workspace_uuid}/user/{user_uuid}/"
	RouteUserGetOne              = "/user/{user_uuid}/"
	RouteWorkspaceInvite         = "/workspace/{workspace_uuid}/invite/"
	RouteActivateUserInWorkspace = "/workspace/{workspace_uuid}/user/{user_uuid}/allow/"
	RouteSuspendUserInWorkspace  = "/workspace/{workspace_uuid}/user/{user_uuid}/suspend/"
	RouteUserOtpGet              = "/user/{user_uuid}/otp/"
	RouteUserOtpPost             = "/user/{user_uuid}/otp/"
	RouteUserOtpDelete           = "/user/{user_uuid}/otp/"

	RouteUserCreate      = "/user/"
	RouteUserSecretGet   = "/user/{user_uuid}/secret/"
	RouteUserSecretPost  = "/user/{user_uuid}/secret/"
	RouteUserGetAll      = "/user/"
	RouteUserDelete      = "/user/{user_uuid}/"
	RouteUserGetAllRoles = "/user/{user_uuid}/role/"

	RouteUserTokenGetOne = "/user/{user_uuid}/user-token/{user_token_uuid}/"

	RouteBackupKeyGetOne = "/workspace/{workspace_uuid}/backup-key/{backup_key_uuid}/"

	RouteServiceUserGetByName = "/workspace/workspace={workspace_name}/service-user/name={service_user_name}/"

	RouteRuleGetAll = "/rule/"

	RouteKiseSecretCreate         = "/workspace/{workspace_uuid}/user/{user_uuid}/kise/key/"
	RouteKiseSecretGetAll         = "/workspace/{workspace_uuid}/user/{user_uuid}/kise/key/"
	RouteKiseSecretDelete         = "/workspace/{workspace_uuid}/user/{user_uuid}/kise/key/{kise_secret_uuid}/"
	RouteRoleGetAll               = "/role/"
	RouteServiceUserGetBindedRole = "/workspace/{workspace_uuid}/role/{role_uuid}/service-user/{service_user_uuid}/"

	RoutePublicKeyVerify = "/public-key/verify/"
	RoutePublicKeyGetOne = "/user/{user_uuid}/public-key/{public_key_uuid}/"
	RouteRoleUpdate      = "/workspace/{workspace_uuid}/role/{role_uuid}/"

	RouteServiceGetAll = "/service/"
	RouteServiceGetOne = "/service/{name}/"

	RouteUserGetOneWorkspace       = "/user/{user_uuid}/workspace/"
	RouteUserGetOneWorkspaceByName = "/user/{user_uuid}/workspace/name={workspace_name}/"
	RouteUserGetOneRoleByName      = "/user/{user_uuid}/workspace/name={workspace_name}/role/name={role_name}/"
	RouteUserGetOneRuleByName      = "/user/{user_uuid}/workspace/name={workspace_name}/rule/name={rule_name}/"
	RouteUserGetAllRules           = "/user/{user_uuid}/rule/"
	RouteUserAppendWorkspace       = "/workspace/{workspace_uuid}/user/{user_uuid}/"
	RouteUserAppendRole            = "/workspace/{workspace_uuid}/role/{role_uuid}/user/{user_uuid}/"
	RouteUserGetBindedRole         = "/workspace/{workspace_uuid}/role/{role_uuid}/user/{user_uuid}/"
	RouteUserUpdate                = "/user/{user_uuid}/"
	RouteUserActivate              = "/user/{user_uuid}/activate/"
	RouteUserSuspend               = "/user/{user_uuid}/suspend/"
	RouteUserTokenIdentify         = "/identify/"
	RouteIdentifyAndAuthorize      = "/identify-and-authz/"

	RouteWorkspaceCreate         = "/workspace/"
	RouteWorkspaceGetOne         = "/workspace/{workspace_uuid}/"
	RouteWorkspaceGetAll         = "/workspace/"
	RouteWorkspaceDelete         = "/workspace/{workspace_uuid}/"
	RouteWorkspaceUserGetByEmail = "/workspace/{workspace_uuid}/user/?email={user_email}"

	RouteGroupUserGetOne        = "/workspace/{workspace_uuid}/group/{group_uuid}/user/{user_uuid}/"
	RouteGroupServiceUserGetOne = "/workspace/{workspace_uuid}/group/{group_uuid}/service-user/{service_user_uuid}/"
	RouteGroupGetByName         = "/user/{user_uuid}/workspace/name={workspace_name}/group/name={group_name}/"
	RouteRoleGetAllGroups       = "/workspace/{workspace_uuid}/role/{role_uuid}/group/"
	RouteGroupAppendRole        = "/workspace/{workspace_uuid}/role/{role_uuid}/group/{group_uuid}/"
	RouteGroupDropRole          = "/workspace/{workspace_uuid}/role/{role_uuid}/group/{group_uuid}/"

	RouteOrganizationGetAll           = "/organization/"
	RouteOrganizationGetOne           = "/organization/{organization_uuid}/"
	RouteOrganizationWorkspacesGetAll = "/organization/{organization_uuid}/workspace/"
	RouteOrganizationWorkspacesGetOne = "/organization/{organization_uuid}/workspace/{workspace_uuid}/"
)
