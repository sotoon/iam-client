package client

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/types"
)

// mockClient implements the Client interface for testing purposes
type mockClient struct {
	Client
}

// UpdateGroup implements the UpdateGroup method for the mock client
func (m *mockClient) UpdateGroup(workspaceUUID, groupUUID uuid.UUID, name, description *string, workspaceInfo *types.WorkspaceUpdateReq) error {
	return nil
}

// BulkAddRolesToGroup implements the BulkAddRolesToGroup method for the mock client
func (m *mockClient) BulkAddRolesToGroup(workspaceUUID, groupUUID uuid.UUID, rolesWithItems []types.RoleWithItems) error {
	return nil
}

// UnbindRoleFromGroup implements the UnbindRoleFromGroup method for the mock client
func (m *mockClient) UnbindRoleFromGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error {
	return nil
}

// BindUserToGroup implements the BindUserToGroup method for the mock client
func (m *mockClient) BindUserToGroup(workspaceUUID, groupUUID, userUUID *uuid.UUID) error {
	return nil
}

// NewMockClient creates a new mock client that wraps an existing client
func NewMockClient(client Client) Client {
	return &mockClient{
		Client: client,
	}
}
