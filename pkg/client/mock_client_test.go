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
func (m *mockClient) UpdateGroup(workspaceUUID, groupUUID *uuid.UUID, name, description string, workspaceInfo ...types.WorkspaceUpdateReq) error {
	return nil
}

// NewMockClient creates a new mock client that wraps an existing client
func NewMockClient(client Client) Client {
	return &mockClient{
		Client: client,
	}
}
