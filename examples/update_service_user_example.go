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
	serviceUserId := "{service_user_id}" // Replace with an actual service user UUID

	// Updated service user details
	updatedName := "updated-service-user"
	updatedDescription := "This is an updated description for the service user"

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

	serviceUserUUID, err := uuid.FromString(serviceUserId)
	if err != nil {
		fmt.Println("invalid service user UUID:", err)
		os.Exit(1)
	}

	updatedServiceUser, err := client.UpdateServiceUser(workspaceUUID, serviceUserUUID, updatedName, updatedDescription)
	if err != nil {
		fmt.Println("error updating service user:", err)
		os.Exit(1)
	}

	fmt.Println("Service User updated successfully!")

	jsonData, err := json.MarshalIndent(updatedServiceUser, "", "  ")
	if err != nil {
		fmt.Println("error marshaling service user data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
