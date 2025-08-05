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

	serviceUsers, err := client.GetAllGroupServiceUserList(&workspaceUUID, &groupUUID)
	if err != nil {
		fmt.Println("error getting group service users:", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d service users in group\n", len(serviceUsers))

	for i, serviceUser := range serviceUsers {
		fmt.Printf("%d. Service User: %s (UUID: %s)\n", i+1, serviceUser.Name, serviceUser.UUID)
		fmt.Printf("   Description: %s\n", serviceUser.Description)
	}

	jsonData, err := json.MarshalIndent(serviceUsers, "", "  ")
	if err != nil {
		fmt.Println("error marshaling service user data:", err)
		os.Exit(1)
	}
	fmt.Println("\nFull service user details:")
	fmt.Println(string(jsonData))
}
