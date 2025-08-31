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
	userUUID := "{user_uuid}"

	publicKeyId := "{public_key_uuid}"

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID, client.DEBUG)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	publicKeyUUID, err := uuid.FromString(publicKeyId)
	if err != nil {
		fmt.Println("invalid public key UUID:", err)
		os.Exit(1)
	}

	err = client.DeleteMyUserPublicKey(&publicKeyUUID)
	if err != nil {
		fmt.Println("error deleting public key:", err)
		os.Exit(1)
	}

	fmt.Println("Public key deleted successfully!")
}
