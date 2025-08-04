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

	users, err := client.GetAllGroupUserList(&workspaceUUID, &groupUUID)
	if err != nil {
		fmt.Println("error getting group users:", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d users in group\n", len(users))

	for i, user := range users {
		fmt.Printf("%d. User: %s %s (UUID: %s)\n", i+1, user.FirstName, user.LastName, user.UUID)
		fmt.Printf("   Email: %s\n", user.Email)
	}

	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("error marshaling user data:", err)
		os.Exit(1)
	}
	fmt.Println("\nFull user details:")
	fmt.Println(string(jsonData))
}
