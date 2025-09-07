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
	IAM_URL := client.BepaURL
	workspaceId := "{your_workspace_uuid}"

	// Role name to create
	roleName := "{your_role_name}"
	roleDescription := "{your_role_description}"

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

	createdRole, err := client.CreateRole(roleName, roleDescription, &workspaceUUID)
	if err != nil {
		fmt.Println("error creating role:", err)
		os.Exit(1)
	}

	fmt.Println("Role created successfully!")

	jsonData, err := json.MarshalIndent(createdRole, "", "  ")
	if err != nil {
		fmt.Println("error marshaling role data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
