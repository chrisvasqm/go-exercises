package greet

import "fmt"

func Greet(name string) string {
	message := fmt.Sprintf("Hello %v, welcome!", name)
	return message
}