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
	groupId := "{group_id}" // Replace with an actual group UUID

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

	groupUUID, err := uuid.FromString(groupId)
	if err != nil {
		fmt.Println("invalid group UUID:", err)
		os.Exit(1)
	}

	newName := "Updated Group Name"
	newDescription := "This is an updated description for the group"
	err = client.UpdateGroup(workspaceUUID, groupUUID, &newName, &newDescription, nil)
	if err != nil {
		fmt.Println("error updating group:", err)
		os.Exit(1)
	}

}
