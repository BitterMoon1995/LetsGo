package main

import "fmt"

func test(n int) {
	if n > 2 {
		fmt.Println(n)
		n--
		test(n)
	}
}

func alterable(i ...int) {
	fmt.Println(i)
}

func main() {
	alterable(1, 2, 3)
}
