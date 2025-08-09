package main

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_id}"
	roleId := "{role_id}"

	// Create a list of user UUIDs to add
	userIds := []string{
		"{user_id_1}",
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

	userUUIDs := make([]uuid.UUID, len(userIds))
	for i, id := range userIds {
		userUUID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("error parsing user UUID %s: %v\n", id, err)
			os.Exit(1)
		}
		userUUIDs[i] = userUUID
	}

	// Bulk add users to role
	err = client.BulkAddUsersToRole(workspaceUUID, roleUUID, userUUIDs)
	if err != nil {
		fmt.Println("error bulk adding users to role:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully bulk added users to the role!")
}
