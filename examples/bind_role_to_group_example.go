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
	groupId := "{group_id}"
	roleId := "{role_id}"

	items := map[string]string{
		"key_one": "value1",
		"key_two": "value2",
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

	groupUUID, err := uuid.FromString(groupId)
	if err != nil {
		fmt.Println("error parsing group UUID:", err)
		os.Exit(1)
	}

	roleUUID, err := uuid.FromString(roleId)
	if err != nil {
		fmt.Println("error parsing role UUID:", err)
		os.Exit(1)
	}

	// Bind role to group with items
	err = client.BindRoleToGroup(&workspaceUUID, &roleUUID, &groupUUID, items)
	if err != nil {
		fmt.Println("error binding role to group:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully bound role to the group!")
}
