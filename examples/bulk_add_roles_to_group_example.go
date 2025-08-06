package main

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
	"github.com/sotoon/iam-client/pkg/types"
)

func main() {
	accessToken := "{access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_id}"
	groupId := "{group_id}"

	// Role UUIDs to add to the group
	roleIds := []string{
		"ca9c0634-3296-46a1-8416-007bdc388dc7",
	}
	items := []map[string]string{ // items if the role requeies items
		{"key_one": "value1"},
		{"key_two": "value2"},
	}

	client, err := client.NewClient(accessToken, IAM_URL, "", "", client.DEBUG)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	workspaceUUID, err := uuid.FromString(workspaceId)
	if err != nil {
		fmt.Println("invalid workspace UUID:", err)
		os.Exit(1)
	}

	groupUUID, err := uuid.FromString(groupId)
	if err != nil {
		fmt.Println("invalid group UUID:", err)
		os.Exit(1)
	}

	// Convert string role IDs to UUID objects
	roleUUIDs := make([]uuid.UUID, len(roleIds))
	for i, roleId := range roleIds {
		roleUUID, err := uuid.FromString(roleId)
		if err != nil {
			fmt.Printf("invalid role UUID %s: %v\n", roleId, err)
			os.Exit(1)
		}
		roleUUIDs[i] = roleUUID
	}

	// Create roles with items
	rolesWithItems := make([]types.RoleWithItems, len(roleUUIDs))
	for i, roleUUID := range roleUUIDs {
		rolesWithItems[i] = types.RoleWithItems{
			RoleUUID: roleUUID.String(),
			Items:    items,
		}
	}

	// Bulk add roles to the group with items
	err = client.BulkAddRolesToGroup(workspaceUUID, groupUUID, rolesWithItems)
	if err != nil {
		fmt.Println("error adding roles to group:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully added roles to the group!")
}
