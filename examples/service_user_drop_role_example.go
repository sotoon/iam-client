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
	serviceUserId := "{service_user_id}"

	// Optional items to include as query parameters
	items := map[string]string{
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

	roleUUID, err := uuid.FromString(roleId)
	if err != nil {
		fmt.Println("error parsing role UUID:", err)
		os.Exit(1)
	}

	serviceUserUUID, err := uuid.FromString(serviceUserId)
	if err != nil {
		fmt.Println("error parsing service user UUID:", err)
		os.Exit(1)
	}

	// Unbind role from service user with items
	err = client.UnbindRoleFromServiceUser(&workspaceUUID, &roleUUID, &serviceUserUUID, items)
	if err != nil {
		fmt.Println("error unbinding role from service user:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully unbound role from the service user!")
}
