package main

import (
	"encoding/json"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{access_token}"
	IAM_URL := client.BepaURL
	workspaceId := "{workspace_id}"
	description := "Test Group Description"
	// Group name to create
	groupName := "Test Group "

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

	createdGroup, err := client.CreateGroup(groupName, description, &workspaceUUID)
	if err != nil {
		fmt.Println("error creating group:", err)
		os.Exit(1)
	}

	fmt.Println("Group created successfully!")

	jsonData, err := json.MarshalIndent(createdGroup, "", "  ")
	if err != nil {
		fmt.Println("error marshaling group data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
