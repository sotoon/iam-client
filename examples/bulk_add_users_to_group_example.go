package main

import (
	"encoding/json"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{your_access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_uuid}"
	groupId := "{group_uuid}"

	userIds := []string{
		"{user_uuid}",
		"{user_uuid}",
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

	userUUIDs := make([]uuid.UUID, len(userIds))
	for i, id := range userIds {
		userUUID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("invalid user UUID %s: %v\n", id, err)
			os.Exit(1)
		}
		userUUIDs[i] = userUUID
	}

	groupUsers, err := client.BulkAddUsersToGroup(workspaceUUID, groupUUID, userUUIDs)
	if err != nil {
		fmt.Println("error adding users to group:", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully added %d users to group %s\n", len(userUUIDs), groupId)
	fmt.Println("\nFull group user details:")
	jsonData, err := json.MarshalIndent(groupUsers, "", "  ")
	if err != nil {
		fmt.Println("error marshaling group user data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
