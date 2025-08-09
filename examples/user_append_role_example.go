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
	userId := "{user_id}"

	// Optional items to include in the request
	items := map[string]string{}

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

	userUUID, err := uuid.FromString(userId)
	if err != nil {
		fmt.Println("error parsing user UUID:", err)
		os.Exit(1)
	}

	// Bind role to user with items
	err = client.BindRoleToUser(&workspaceUUID, &roleUUID, &userUUID, items)
	if err != nil {
		fmt.Println("error binding role to user:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully bound role to the user!")
}
