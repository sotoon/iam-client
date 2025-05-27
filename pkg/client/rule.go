package client

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) CreateRule(ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*types.Rule, error) {
	ruleRequest := &types.RuleReq{
		Name:    ruleName,
		Actions: ruleActions,
		Object:  object,
		Deny:    deny,
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleCreate), replaceDict)

	createdRule := &types.Rule{}
	if err := c.Do(http.MethodPost, apiURL, 0, ruleRequest, createdRule); err != nil {
		return nil, err
	}
	return createdRule, nil
}

func (c *iamClient) UpdateRule(ruleUUID *uuid.UUID, ruleName string, workspaceUUID *uuid.UUID, ruleActions []string, object string, deny bool) (*types.Rule, error) {
	ruleRequest := &types.RuleReq{
		Name:    ruleName,
		Actions: ruleActions,
		Object:  object,
		Deny:    deny,
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleUpdate), replaceDict)

	updatedRule := &types.Rule{}
	if err := c.Do(http.MethodPatch, apiURL, 0, ruleRequest, updatedRule); err != nil {
		return nil, err
	}
	return updatedRule, nil
}

func (c *iamClient) DeleteRule(ruleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) GetRuleRoles(ruleUUID, workspaceUUID *uuid.UUID) ([]*types.Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetAllRoles), replaceDict)

	roles := []*types.Role{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *iamClient) BindRuleToRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleAppendRule), replaceDict)
	err := c.Do(http.MethodPost, apiURL, 0, nil, nil)
	return err
}

func (c *iamClient) UnbindRuleFromRole(roleUUID, ruleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleDropRule), replaceDict)
	err := c.Do(http.MethodDelete, apiURL, 0, nil, nil)
	return err
}

func (c *iamClient) GetRule(ruleUUID, workspaceUUID *uuid.UUID) (*types.Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		ruleUUIDPlaceholder:      ruleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetOne), replaceDict)

	rule := &types.Rule{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func (c *iamClient) GetRuleByName(ruleName, workspaceName string) (*types.Rule, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: workspaceName,
		ruleNamePlaceholder:      ruleName,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneRuleByName), replaceDict)
	rule := &types.Rule{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func (c *iamClient) GetAllRules() ([]*types.Rule, error) {
	replaceDict := map[string]string{}
	apiURL := substringReplace(trimURLSlash(routes.RouteRuleGetAll), replaceDict)

	rules := []*types.Rule{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *iamClient) GetAllUserRules(userUUID *uuid.UUID) ([]*types.Rule, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllRules), replaceDict)

	rules := []*types.Rule{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}
