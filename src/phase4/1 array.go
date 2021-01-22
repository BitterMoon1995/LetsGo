package main

import "fmt"

/*
数组的长度也是它的类型的一部分
*/
func declareArray() {
	var arr [4]int
	var arr2 []int
	var strArr []string
	fmt.Printf("%T %T %T", arr, arr2, strArr)
}

func initArray() {
	var arr1 [3]int
	//arr1[3] = 5
	fmt.Println(arr1)

	var arr2 = [3]int{6, 1, 8}
	fmt.Println(arr2)

	/*短变量声明数组*/
	arr3 := [4]string{"stop", "hate", "me", "nigger"}
	fmt.Println(arr3)

	/*长度用...声明，让编译器自行判断*/
	arr4 := [...]string{"stop", "hate", "me", "cause", "I`m", "beautiful", "nigger"}
	fmt.Println(len(arr4))
	//arr4[8] = "nigga"

	/*初始化数组时可以指定索引值*/
	arr5 := [...]int{0: 1, 3: 8, 5: 12, 7: 32}
	fmt.Println(arr5)
}

func addrOfArray() {
	arr := [...]int{1, 2, 3}
	arrP := &arr
	fmt.Println(arrP, &arr[0], &arr[1])
}

func arrUseValueCopy() {
	arr := [3]int{1, 2, 3}
	//arrUseValueCopyT(arr)
	arrUseValueCopyT2(&arr)
	fmt.Println(arr)
}
func arrUseValueCopyT(arr [3]int) {
	arr[0] = 6324
}
func arrUseValueCopyT2(arrP *[3]int) {
	(*arrP)[0] = 4396
}

func traverseByRange() {
	arr := [...]int{0: 1, 3: 8, 5: 12}
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

func twoDimensional() {
	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr)

	/*二维数组就是两个连在一起的数组*/
	fmt.Println(&arr[0][0])
	fmt.Println(&arr[1][0])

}

func main() {
	twoDimensional()
}
