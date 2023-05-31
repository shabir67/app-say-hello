package main

import (
	"fmt"

	gosayhello "github.com/shabir67/go-say-hello"
)

func main() {
	hello := gosayhello.SayHello()
	fmt.Println(hello, "Joko")
}
