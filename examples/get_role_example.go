package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	accessToken := "{access_token}"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_uuid}"
	roleId := "{role_uuid}"

	// Parse UUIDs
	workspaceUUID, err := uuid.FromString(workspaceId)
	if err != nil {
		log.Fatalf("error parsing workspace UUID: %v", err)
	}

	roleUUID, err := uuid.FromString(roleId)
	if err != nil {
		log.Fatalf("error parsing role UUID: %v", err)
	}

	// Create IAM client
	iamClient, err := client.NewClient(accessToken, IAM_URL, "", "", client.DEBUG)
	if err != nil {
		log.Fatalf("error creating IAM client: %v", err)
	}

	// Get role details
	role, err := iamClient.GetRole(&roleUUID, &workspaceUUID)
	if err != nil {
		log.Fatalf("error getting role: %v", err)
	}

	// Pretty print the role details
	roleJSON, err := json.MarshalIndent(role, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling role to JSON: %v", err)
	}

	fmt.Println("Role Details:")
	fmt.Println(string(roleJSON))
}
