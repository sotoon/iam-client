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

	client, err := client.NewClient(accessToken, IAM_URL, "", "{user_uuid}")
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	workspaceUUID, err := uuid.FromString(workspaceId)
	if err != nil {
		fmt.Println("invalid workspace UUID:", err)
		os.Exit(1)
	}

	serviceUsers, err := client.GetWorkspaceServiceUserList(workspaceUUID)
	if err != nil {
		fmt.Println("error getting service user details:", err)
		os.Exit(1)
	}

	fmt.Println("Service Users:")
	for i, serviceUser := range serviceUsers {
		fmt.Println(i)
		res, _ := json.Marshal(serviceUser)
		fmt.Println(string(res))
		fmt.Println("   ---")
	}
}
