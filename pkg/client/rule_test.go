package client

import (
	"net/http"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
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

	config := testutils.TestConfig{
		Object:           &rule,
		Params:           []interface{}{ruleName, &workspaceUUID, ruleActions, object, deny},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/rule/`),
		ClientMethodName: "CreateRule",
	}
	testutils.DoTestCreateAPI(t, config)
}
func TestGetRuleRoles(t *testing.T) {
	var roles []types.Role
	ruleUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &roles,
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.*)/rule/(.*)/role/$`),
		ClientMethodName: "GetRuleRoles",
		Params:           []interface{}{&ruleUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &ruleUUID},
	}
	testutils.DoTestListingAPI(t, config)
}
func TestBindRuleToRole(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&roleUUID, &ruleUUID, &workspaceUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &ruleUUID},

		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/(.+)/rule/(.+)/`),
		ClientMethodName: "BindRuleToRole",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindRuleFromRole(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	ruleUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&roleUUID, &ruleUUID, &workspaceUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &ruleUUID},

		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/(.+)/rule/(.+)/`),
		ClientMethodName: "UnbindRuleFromRole",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestGetRule(t *testing.T) {
	var object types.Rule
	rule := uuid.NewV4()
	workspace := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&rule, &workspace},
		ParamsInURL:      []interface{}{&workspace, &rule},
		ParamNames:       []string{"Name", "Workspace"},
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.+)/rule/(.+)/$`),
		ClientMethodName: "GetRule",
	}

	testutils.DoTestReadAPI(t, config)
}
func TestGetRuleByName(t *testing.T) {
	var object types.Rule
	var ruleName, workspaceName string
	faker.FakeData(&ruleName)
	faker.FakeData(&workspaceName)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{ruleName, workspaceName},
		ParamsInURL:      []interface{}{workspaceName, ruleName},
		ParamNames:       []string{"Name", "Workspace"},
		URLregexp:        regexp.MustCompile(`^/api/v1/user/.+/workspace/name=(.+)/rule/name=(.+)/$`),
		ClientMethodName: "GetRuleByName",
	}

	testutils.DoTestReadAPI(t, config)
}
func TestGetAllRules(t *testing.T) {
	rules := []types.Rule{}

	config := testutils.TestConfig{
		Object:           &rules,
		URLregexp:        regexp.MustCompile(`^/api/v1/rule/$`),
		ClientMethodName: "GetAllRules",
	}
	testutils.DoTestListingAPI(t, config)
}

func TestGetAllUserRules(t *testing.T) {
	rules := []types.Rule{}
	userUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &rules,
		URLregexp:        regexp.MustCompile(`^/api/v1/user/(.+)/rule/$`),
		ClientMethodName: "GetAllUserRules",
		Params:           []interface{}{&userUUID},
		ParamsInURL:      []interface{}{&userUUID},
	}
	testutils.DoTestListingAPI(t, config)
}
