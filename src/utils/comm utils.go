package utils

import "fmt"

var F = func(a int, b int) int {
	return a * b
}

func init() {

	fmt.Println("init() of comm utils")
	Name = "Nigga"
	Age = 55
}
