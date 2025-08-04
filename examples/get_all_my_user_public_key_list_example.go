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

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	publicKeys, err := client.GetAllMyUserPublicKeyList()
	if err != nil {
		fmt.Println("error getting user public keys:", err)
		os.Exit(1)
	}

	fmt.Println("User Public Keys:")
	if len(publicKeys) == 0 {
		fmt.Println("No public keys found for this user")
	} else {
		fmt.Printf("Found %d public keys\n", len(publicKeys))

		jsonData, err := json.MarshalIndent(publicKeys, "", "  ")
		if err != nil {
			fmt.Println("error marshaling public key data:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	}
}
