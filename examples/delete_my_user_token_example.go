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
	tokenId := "{token_uuid}"

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	tokenUUID, err := uuid.FromString(tokenId)
	if err != nil {
		fmt.Println("invalid token UUID:", err)
		os.Exit(1)
	}

	err = client.DeleteMyUserToken(&tokenUUID)
	if err != nil {
		fmt.Println("error deleting user token:", err)
		os.Exit(1)
	}

	fmt.Println("User token deleted successfully!")
}
