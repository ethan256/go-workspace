package main

import (
	"fmt"

	"github.com/ethan256/util"
)

func main() {
	msg := "test workspace!"
	resp := modea.Consume(msg)
	fmt.Println(resp)
}
