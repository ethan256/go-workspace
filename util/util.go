package util

import "fmt"

func Consume(msg string) string {
	return fmt.Sprintf("consume msg: %s", msg)
}
