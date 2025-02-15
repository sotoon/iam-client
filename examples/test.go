package main

import (
	"fmt"
	"time"

	"git.platform.sotoon.ir/iam/golang-bepa-client/pkg/client"
)

func main() {
	urls := []string{"https://afra.bepa.sotoon.ir", "https://neda.bepa.sotoon.ir"}
	c, _ := client.NewReliableClient("", urls, "", "", 5*time.Second)
	t, e := c.CreateUserTokenByCreds("foo@bar.ir", "__")

	fmt.Println(e)
	fmt.Println(t)
}
