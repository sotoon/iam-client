package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{your_access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	userUUID := "{user_uuid}"

	title := "{title_for_public_key}"
	keyType := "{key_type_of_key}"
	key := "{key}"

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID, client.DEBUG)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	createdPublicKey, err := client.CreateMyUserPublicKey(title, keyType, key)
	if err != nil {
		fmt.Println("error creating public key:", err)
		os.Exit(1)
	}

	fmt.Println("Public Key Created Successfully!")

	jsonData, err := json.MarshalIndent(createdPublicKey, "", "  ")
	if err != nil {
		fmt.Println("error marshaling public key data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
