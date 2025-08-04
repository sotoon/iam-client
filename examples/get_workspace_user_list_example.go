package main

import (
	"encoding/json"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "8b7d01d141038287c977995fc3048c4f9a5a00c5fce0889876e834e478d1af0b14f7269a3e3767e93fe8f142349fe9c662ad070e22bd18d9b9ca4c17ffbdc714"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "d750a6ac-65bf-498c-80cf-a87d52c911a1"

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
