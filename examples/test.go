package main

import (
	"fmt"
	"time"

	"github.com/sotoon/iam-client/pkg/client"
)

func main() {
	urls := []string{"https://afra.bepa.sotoon.ir", "https://neda.bepa.sotoon.ir"}
	c, _ := client.NewReliableClient("", urls, "", "", 5*time.Second)
	t, e := c.CreateUserTokenByCreds("foo@bar.ir", "__")

	fmt.Println(e)
	fmt.Println(t)
}
