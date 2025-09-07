package main

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{your_access_token}"
	IAM_URL := client.BepaURL
	workspaceId := "{your_workspace_uuid}"
	roleId := "{your_role_uuid}"

	// Create a list of rule UUIDs to add
	ruleIds := []string{
		"{rule_uuid_1}",
		"{rule_uuid_2}",
		// Add more rule UUIDs as needed
	}

	client, err := client.NewClient(accessToken, IAM_URL, "", "", client.DEBUG)
	if err != nil {
		fmt.Println("error creating client:", err)
		os.Exit(1)
	}

	// Parse UUIDs
	workspaceUUID, err := uuid.FromString(workspaceId)
	if err != nil {
		fmt.Println("error parsing workspace UUID:", err)
		os.Exit(1)
	}

	roleUUID, err := uuid.FromString(roleId)
	if err != nil {
		fmt.Println("error parsing role UUID:", err)
		os.Exit(1)
	}

	ruleUUIDs := make([]uuid.UUID, len(ruleIds))
	for i, id := range ruleIds {
		ruleUUID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("error parsing rule UUID %s: %v\n", id, err)
			os.Exit(1)
		}
		ruleUUIDs[i] = ruleUUID
	}

	// Bulk add rules to role
	err = client.BulkAddRulesToRole(workspaceUUID, roleUUID, ruleUUIDs)
	if err != nil {
		fmt.Println("error bulk adding rules to role:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully bulk added rules to the role!")
}
