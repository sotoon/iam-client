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

	publicKeyId := "{public_key_uuid}"

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

	publicKeyUUID, err := uuid.FromString(publicKeyId)
	if err != nil {
		fmt.Println("invalid public key UUID:", err)
		os.Exit(1)
	}

	err = client.DeleteServiceUserPublicKey(workspaceUUID, serviceUserUUID, publicKeyUUID)
	if err != nil {
		fmt.Println("error deleting service user public key:", err)
		os.Exit(1)
	}

	fmt.Println("Service user public key deleted successfully!")
}
