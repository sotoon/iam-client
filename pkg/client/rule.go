package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Rule struct {
	UUID          *uuid.UUID     `json:"uuid"`
	WorkspaceUUID *uuid.UUID     `json:"workspace"`
	Name          string         `json:"name"`
	Actions       pq.StringArray `json:"actions" validate:"required,gte=1"`
	Object        string         `json:"object" validate:"required"`
	Deny          bool           `json:"deny"`
}

func (c *bepaClient) CreateRule(ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, obejct string, deny bool) (*Rule, error) {
	ruleRequest := &types.RuleReq{
		Name:    ruleName,
		Actions: ruleActions,
		Object:  obejct,
		Deny:    deny,
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleCreate), replaceDict)

	createdRule := &Rule{}
	if err := c.Do(http.MethodPost, apiURL, ruleRequest, createdRule); err != nil {
		return nil, err
	}
	return createdRule, nil
}

func (c *bepaClient) GetRuleRoles(ruleUUID, workspaceUUID *uuid.UUID) ([]*Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetAllRoles), replaceDict)

	roles := []*Role{}
	if err := c.Do(http.MethodGet, apiURL, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *bepaClient) BindRuleToRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleAppendRule), replaceDict)
	err := c.Do(http.MethodPost, apiURL, nil, nil)
	return err
}

func (c *bepaClient) UnbindRuleFromRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleDropRule), replaceDict)
	err := c.Do(http.MethodDelete, apiURL, nil, nil)
	return err
}

func (c *bepaClient) GetRule(ruleUUID, workspaceUUID *uuid.UUID) (*Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetOne), replaceDict)

	rule := &Rule{}
	if err := c.Do(http.MethodGet, apiURL, nil, rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func (c *bepaClient) GetRuleByName(ruleName, workspaceName string) (*Rule, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: workspaceName,
		ruleNamePlaceholder:      ruleName,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneRuleByName), replaceDict)
	rule := &Rule{}
	if err := c.Do(http.MethodGet, apiURL, nil, rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func (c *bepaClient) GetAllRules() ([]*Rule, error) {
	replaceDict := map[string]string{}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetAll), replaceDict)

	rules := []*Rule{}
	if err := c.Do(http.MethodGet, apiURL, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *bepaClient) GetAllUserRules(userUUID *uuid.UUID) ([]*Rule, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllRules), replaceDict)

	rules := []*Rule{}
	if err := c.Do(http.MethodGet, apiURL, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}
