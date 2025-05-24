package client

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/sotoon/iam-client/pkg/types"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestCreateRule(t *testing.T) {
	var rule types.Rule
	var ruleName, object string
	var ruleActions []string
	var deny bool

	workspaceUUID := uuid.NewV4()
	faker.FakeData(&ruleName)
	faker.FakeData(&ruleActions)
	faker.FakeData(&deny)

	config := TestConfig{
		Object:           &rule,
		Params:           []interface{}{ruleName, &workspaceUUID, ruleActions, object, deny},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/rule/`),
		ClientMethodName: "CreateRule",
	}
	DoTestCreateAPI(t, config)
}

func TestUpdateRule(t *testing.T) {
	var rule types.Rule
	var ruleName, object string
	var ruleActions []string
	var deny bool

	workspaceUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()
	faker.FakeData(&ruleName)
	faker.FakeData(&ruleActions)
	faker.FakeData(&deny)
	faker.FakeData(&object)

	config := TestConfig{
		Object:           &rule,
		Params:           []interface{}{&ruleUUID, ruleName, &workspaceUUID, ruleActions, object, deny},
		ParamNames:       []string{"", "Name", "", "Actions", "Object", "Deny"},
		ParamsInURL:      []interface{}{&workspaceUUID, &ruleUUID},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/rule/(.+)/`),
		ClientMethodName: "UpdateRule",
	}
	DoTestUpdateAPI(t, config, http.MethodPatch)
}

func TestDeleteRule(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()

	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/rule/(.+)/`),
		ClientMethodName: "DeleteRule",
		Params:           []interface{}{&workspaceUUID, &ruleUUID},
		ParamsInURL:      []interface{}{&ruleUUID, &workspaceUUID},
	}
	DoTestDeleteAPI(t, conf)
}

func TestGetRuleRoles(t *testing.T) {
	var roles []types.Role
	ruleUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	config := TestConfig{
		Object:           &roles,
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.*)/rule/(.*)/role/$`),
		ClientMethodName: "GetRuleRoles",
		Params:           []interface{}{&ruleUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &ruleUUID},
	}
	DoTestListingAPI(t, config)
}
func TestBindRuleToRole(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := TestConfig{
		Params:      []interface{}{&roleUUID, &ruleUUID, &workspaceUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &ruleUUID},

		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/(.+)/rule/(.+)/`),
		ClientMethodName: "BindRuleToRole",
	}
	DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindRuleFromRole(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := TestConfig{
		Params:      []interface{}{&roleUUID, &ruleUUID, &workspaceUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &ruleUUID},

		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/(.+)/rule/(.+)/`),
		ClientMethodName: "UnbindRuleFromRole",
	}
	DoTestDeleteAPI(t, config)
}

func TestGetRule(t *testing.T) {
	var object types.Rule
	rule := uuid.NewV4()
	workspace := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&rule, &workspace},
		ParamsInURL:      []interface{}{&workspace, &rule},
		ParamNames:       []string{"Name", "Workspace"},
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.+)/rule/(.+)/$`),
		ClientMethodName: "GetRule",
	}

	DoTestReadAPI(t, config)
}
func TestGetRuleByName(t *testing.T) {
	var object types.Rule
	var ruleName, workspaceName string
	faker.FakeData(&ruleName)
	faker.FakeData(&workspaceName)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{ruleName, workspaceName},
		ParamsInURL:      []interface{}{workspaceName, ruleName},
		ParamNames:       []string{"Name", "Workspace"},
		URLregexp:        regexp.MustCompile(`^/api/v1/user/.+/workspace/name=(.+)/rule/name=(.+)/$`),
		ClientMethodName: "GetRuleByName",
	}

	DoTestReadAPI(t, config)
}
func TestGetAllRules(t *testing.T) {
	rules := []types.Rule{}

	config := TestConfig{
		Object:           &rules,
		URLregexp:        regexp.MustCompile(`^/api/v1/rule/$`),
		ClientMethodName: "GetAllRules",
	}
	DoTestListingAPI(t, config)
}

func TestGetAllUserRules(t *testing.T) {
	rules := []types.Rule{}
	userUUID := uuid.NewV4()
	config := TestConfig{
		Object:           &rules,
		URLregexp:        regexp.MustCompile(`^/api/v1/user/(.+)/rule/$`),
		ClientMethodName: "GetAllUserRules",
		Params:           []interface{}{&userUUID},
		ParamsInURL:      []interface{}{&userUUID},
	}
	DoTestListingAPI(t, config)
}
