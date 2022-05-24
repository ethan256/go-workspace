package main

import (
	"fmt"

	"github.com/ethan256/util"
)

func main() {
	msg := "test workspace!"
	resp := util.Consume(msg)
	fmt.Println(resp)
}
