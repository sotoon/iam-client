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

	// Create a list of service user UUIDs to add
	serviceUserIds := []string{
		"{service_user_id_1}",
		"{service_user_id_2}",
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

	serviceUserUUIDs := make([]uuid.UUID, len(serviceUserIds))
	for i, id := range serviceUserIds {
		serviceUserUUID, err := uuid.FromString(id)
		if err != nil {
			fmt.Printf("error parsing service user UUID %s: %v\n", id, err)
			os.Exit(1)
		}
		serviceUserUUIDs[i] = serviceUserUUID
	}

	// Bulk add service users to role
	err = client.BulkAddServiceUsersToRole(workspaceUUID, roleUUID, serviceUserUUIDs)
	if err != nil {
		fmt.Println("error bulk adding service users to role:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully bulk added service users to the role!")
}
