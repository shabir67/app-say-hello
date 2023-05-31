package main

import (
	"fmt"

	gosayhello "github.com/shabir67/go-say-hello/v2"
)

func main() {
	hello := gosayhello.SayHello("Joko")
	fmt.Println(hello)
}
