package main

import (
	"fmt"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/client"
)

func main() {
	var c client.Client
	c, _ = client.NewClient("", "https://bepa.cafebazaar.cloud", "", "")
	t, e := c.CreateUserTokenByCreds("foo@bar.ir", "__")

	fmt.Println(e)
	fmt.Println(t)
}
