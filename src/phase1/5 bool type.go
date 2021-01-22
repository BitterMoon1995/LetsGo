package main

import "fmt"

func defaultTrue() {
	var flag bool
	fmt.Println(flag)
}

/*
bool类型无法参与数值计算，无法与其他类型相互转换
*/
func selfClosing() {
	var flag bool
	if !flag {
		fmt.Println("nigger")
	}

	var x int = 1
	//if x {
	//	fmt.Println("nigger")
	//}

	var f float64 = 13.2
	x = int(f)
	//x = int(flag)
	fmt.Println(x)
}

func main() {
	selfClosing()
}
