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

	client, err := client.NewClient(accessToken, IAM_URL, "", userUUID, client.DEBUG)
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	userTokens, err := client.GetAllMyUserTokenList()
	if err != nil {
		fmt.Println("error getting user tokens:", err)
		os.Exit(1)
	}

	fmt.Println("User Tokens:")
	if userTokens == nil || len(*userTokens) == 0 {
		fmt.Println("No tokens found for this user")
	} else {
		fmt.Printf("Found %d tokens\n", len(*userTokens))

		jsonData, err := json.MarshalIndent(userTokens, "", "  ")
		if err != nil {
			fmt.Println("error marshaling token data:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	}
}
