package main

import (
	"fmt"
        "golang.org/x/net/context"
)

func main () {
	fmt.Println ("Package Name: my package")
        type Context = context.Context
        type CancelFunc = context.CancelFunc
}
