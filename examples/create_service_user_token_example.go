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
	serviceUserId := "{serviceuser_uuid}"

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

	// Create a new token for the service user
	newToken, err := client.CreateServiceUserToken(&serviceUserUUID, &workspaceUUID)
	if err != nil {
		fmt.Println("error creating service user token:", err)
		os.Exit(1)
	}

	// Display the newly created token
	fmt.Println("New Service User Token Created:")
	jsonData, err := json.MarshalIndent(newToken, "", "  ")
	if err != nil {
		fmt.Println("error marshaling token data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
