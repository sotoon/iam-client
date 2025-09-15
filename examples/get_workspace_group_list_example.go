package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/client"
	"github.com/sotoon/iam-client/pkg/client/interceptor"
)

func main() {

	var wg sync.WaitGroup
	count_done := 0

	accessToken := "{access_token}"
	IAM_URL := client.GatewayURL
	workspaceId := "{workspace_id}"

	client, err := client.NewClient(accessToken,
		IAM_URL,
		"",
		"",
		client.INFO,
		client.OptionWithInterceptor([]interceptor.ClientInterceptor{
			interceptor.NewCircuitBreakerInterceptor(interceptor.CircuteBreakerForJust429, false),
		}))
	if err != nil {
		fmt.Println("cannot create client:", err)
		os.Exit(1)
	}

	client.AddInterceptor(
		interceptor.NewRetryInterceptor(
			client,
			interceptor.NewRetryInterceptor_ExponentialBackoff(time.Second, time.Second*10),
			interceptor.NewRetryInterceptor_RetryDeciderAll(10),
		))

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {

			workspaceUUID, err := uuid.FromString(workspaceId)
			if err != nil {
				fmt.Println("invalid workspace UUID:", err)
				os.Exit(1)
			}

			groups, err := client.GetWorkspaceGroupList(workspaceUUID)
			if err != nil {
				fmt.Println("error getting workspace groups:", err)
				os.Exit(1)

			}
			fmt.Printf("Found %d groups in workspace\n", len(groups))

			for i, group := range groups {
				fmt.Printf("%d. Group: %s (UUID: %s)\n", i+1, group.Name, group.UUID)
			}

			fmt.Println("\nFull group details:", i)
			//fmt.Println(string(jsonData))
			count_done++
			fmt.Println("count_done:", count_done)
			wg.Done()
		}(i)

	}
	wg.Wait()
}
