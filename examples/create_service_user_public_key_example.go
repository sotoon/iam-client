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

	keyName := "{key_name_of_key}"
	publicKeyContent := "{public_key_content}"

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

	createdKey, err := client.CreateServiceUserPublicKey(workspaceUUID, serviceUserUUID, keyName, publicKeyContent)
	if err != nil {
		fmt.Println("error creating service user public key:", err)
		os.Exit(1)
	}

	fmt.Println("Service User Public Key Created:")
	jsonData, err := json.MarshalIndent(createdKey, "", "  ")
	if err != nil {
		fmt.Println("error marshaling public key data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
