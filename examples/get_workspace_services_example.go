package main

import (
	"fmt"
	"github.com/sotoon/iam-client/pkg/client"
	"os"
	uuid "github.com/satori/go.uuid"
)

func main() {
	accessToken := "{accessToken}"
	IAM_URL := "https://bepa.sotoon.ir"
	workspaceId := "{workspace_uuid}"

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

	services, err := client.GetWorkspaceServices(workspaceUUID)
	if err != nil {
		fmt.Println("error getting workspace services:", err)
		os.Exit(1)
	}

	fmt.Println("Workspace Services:")
	for _, s := range services {
		fmt.Printf("- %s\n", s.Name)
	}
}
