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
	description := "description of service user"

	// Service user name to create
	serviceUserName := "{service_user_name}"

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

	createdServiceUser, err := client.CreateServiceUser(serviceUserName, description, &workspaceUUID)
	if err != nil {
		fmt.Println("error creating service user:", err)
		os.Exit(1)
	}

	fmt.Println("Service User created successfully!")

	jsonData, err := json.MarshalIndent(createdServiceUser, "", "  ")
	if err != nil {
		fmt.Println("error marshaling service user data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
