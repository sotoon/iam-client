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
	userId := "{user_uuid}"

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

	userUUID, err := uuid.FromString(userId)
	if err != nil {
		fmt.Println("invalid user UUID:", err)
		os.Exit(1)
	}

	user, err := client.GetWorkspaceUserDetail(workspaceUUID, userUUID)
	if err != nil {
		fmt.Println("error getting workspace user detail:", err)
		os.Exit(1)
	}

	fmt.Println("Workspace User Detail:")
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("error marshaling user data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
