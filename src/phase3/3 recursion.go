package main

import "fmt"

func test(n int) {
	if n > 2 {
		fmt.Println(n)
		n--
		test(n)
	}
}

func main() {
	test(10)
}
