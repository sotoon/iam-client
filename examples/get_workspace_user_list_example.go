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
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_id}"

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

	users, err := client.GetWorkspaceUserList(workspaceUUID)
	if err != nil {
		fmt.Println("error getting workspace users:", err)
		os.Exit(1)
	}

	fmt.Println("Workspace Users:")
	if len(users) == 0 {
		fmt.Println("No users found in this workspace")
	} else {
		fmt.Printf("Found %d users in the workspace\n", len(users))

		if len(users) > 0 {
			fmt.Println("\nExample User (Full Details):")
			jsonData, err := json.MarshalIndent(users[0], "", "  ")
			if err != nil {
				fmt.Println("error marshaling user data:", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonData))
		}
	}
}
