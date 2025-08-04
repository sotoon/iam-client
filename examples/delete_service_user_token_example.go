package main

import (
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
	tokenId := "{token_uuid}"

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

	tokenUUID, err := uuid.FromString(tokenId)
	if err != nil {
		fmt.Println("invalid token UUID:", err)
		os.Exit(1)
	}

	err = client.DeleteServiceUserToken(&serviceUserUUID, &workspaceUUID, &tokenUUID)
	if err != nil {
		fmt.Println("error deleting service user token:", err)
		os.Exit(1)
	}

	fmt.Println("Service user token deleted successfully!")
}
