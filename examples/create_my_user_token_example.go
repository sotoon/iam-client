package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{your_access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	userUUID := "{user_uuid}"
	name := "testing token"
	expiresAt := time.Now().Add(24 * time.Hour)

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID, client.DEBUG)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	createdToken, err := client.CreateMyUserToken(name, &expiresAt)
	if err != nil {
		fmt.Println("error creating user token:", err)
		os.Exit(1)
	}

	fmt.Println("User Token Created:")
	jsonData, err := json.MarshalIndent(createdToken, "", "  ")
	if err != nil {
		fmt.Println("error marshaling token data:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
