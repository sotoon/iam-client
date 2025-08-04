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

	serviceUserIds := []string{
		"{service_user_uuid}",
	}

	client, err := client.NewClient(accessToken, IAM_URL, "", "")
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

	serviceUserUUIDs := make([]uuid.UUID, len(serviceUserIds))
	for i, id := range serviceUserIds {
		serviceUserUUID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("invalid service user UUID %s: %v\n", id, err)
			os.Exit(1)
		}
		serviceUserUUIDs[i] = serviceUserUUID
	}

	groupServiceUsers, err := client.BulkAddServiceUsersToGroup(workspaceUUID, groupUUID, serviceUserUUIDs)
	if err != nil {
		fmt.Println("error adding service users to group:", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully added %d service users to group %s\n", len(serviceUserUUIDs), groupId)
	fmt.Println("\nFull group service user details:")
	jsonData, err := json.MarshalIndent(groupServiceUsers, "", "  ")
	if err != nil {
		fmt.Println("error marshaling group service user data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
