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
	userId := "{user_id}" // Replace with actual user UUID

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

	userUUID, err := uuid.FromString(userId)
	if err != nil {
		fmt.Println("invalid user UUID:", err)
		os.Exit(1)
	}

	err = client.UnbindUserFromGroup(&workspaceUUID, &groupUUID, &userUUID)
	if err != nil {
		fmt.Println("error unbinding user from group:", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully unbound user %s from group %s\n", userId, groupId)
}
