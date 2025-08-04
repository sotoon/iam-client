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

	client, err := client.NewClient(accessToken, IAM_URL, "", "")
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

	serviceUser, err := client.GetWorkspaceServiceUserDetail(workspaceUUID, serviceUserUUID)
	if err != nil {
		fmt.Println("error getting service user details:", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(serviceUser, "", "  ")
	if err != nil {
		fmt.Println("error marshaling service user data:", err)
		os.Exit(1)
	}

	fmt.Println("Service User Details:")
	fmt.Println(string(jsonData))
}
