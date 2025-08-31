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

	roles, err := client.GetWorkspaceGroupRoleList(workspaceUUID, groupUUID)
	if err != nil {
		fmt.Println("error getting group roles:", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d roles for group\n", len(roles))

	jsonData, err := json.MarshalIndent(roles, "", "  ")
	if err != nil {
		fmt.Println("error marshaling role data:", err)
		os.Exit(1)
	}
	fmt.Println("\nFull role details:")
	fmt.Println(string(jsonData))
}
