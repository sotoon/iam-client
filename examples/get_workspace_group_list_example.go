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

	groups, err := client.GetWorkspaceGroupList(workspaceUUID)
	if err != nil {
		fmt.Println("error getting workspace groups:", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d groups in workspace\n", len(groups))

	for i, group := range groups {
		fmt.Printf("%d. Group: %s (UUID: %s)\n", i+1, group.Name, group.UUID)
	}

	jsonData, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		fmt.Println("error marshaling group data:", err)
		os.Exit(1)
	}
	fmt.Println("\nFull group details:")
	fmt.Println(string(jsonData))
}
