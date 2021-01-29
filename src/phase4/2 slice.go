package main

import (
	"fmt"
)

/*
通过[n:m]分切引用一个已经存在的数组
*/
func declare() {
	/*一.var声明，值为nil，类型[]int切片（不是指针！），未分配空间*/
	var slice0 []int
	fmt.Println(slice0)

	/*二.带初值声明，和数组的区别就是不写长度*/
	slice := []int{1, 2, 3}
	/*cap函数获取容器容量；切片的容量与长度都是动态改变的*/
	fmt.Println("长度：", len(slice), "容量：", cap(slice))

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	/*三.截取声明，这里取下标3到下标4，不包括下标4*/
	slice = arr[3:4]
	fmt.Println(slice)
	fmt.Println("长度：", len(slice), "容量：", cap(slice))
}

/*
slice是一个结构体，包含所引用首元素地址、长度、容量
*/
func ramModel() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice := arr[2:5]
	fmt.Println("数组下标2元素地址：", &arr[2], "切片下标0元素地址：", &slice[0])
	fmt.Println(len(slice), cap(slice))
	slice[0] = 8848
	fmt.Println("修改后数组：", arr)
}

/*
第四种声明方式：通过make函数声明，同时指定类型和长度。也可以再指定一个容量，容量要大于长度
此时会新创建一个外部不可见的数组，仅由slice进行底层维护
*/
func declareSlice3() {
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
}

/*
第五种声明方式：直接指定具体数组
*/
func declareSlice4() {
	slice := [5]int{}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
}

/*
和数组遍历方式相同
*/
func traverseOfSlice() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for index, element := range slice {
		fmt.Println(index, element)
	}
}

func detailsOfSliceArray() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	/*截取时同样不能超过数组上限*/
	//slice := array[0:6]

	/*截取的上下限可以简写*/
	slice0 := array[:3]
	slice1 := array[3:]
	slice2 := array[:]

	/*切片定义完后，还不能使用，要引用到一个容器才有效*/
	slice := make([]int, 2)
	slice = array[:]
	/* 切片还可以继续截取 */
	sliceOfSlice := slice[:2]

	fmt.Println(slice0)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice)
	fmt.Println(sliceOfSlice)
}

/*
使用append对切片进行动态内容追加
*/
func appendFuncOfSlice() {
	slice := []int{4, 3, 9, 6}
	/*必须接收*/
	appendedSlice := append(slice, 7, 7, 7)
	fmt.Println(appendedSlice)
	/*并不会影响原切片*/
	fmt.Println(slice)
}

/*
append会引起切片的动态扩容，使用cap()内置函数查看当前容器容量
*/
func capacityOfSlice() {
	slice := []int{0}
	for i := 0; i < 10; i++ {
		fmt.Println(cap(slice))
		slice = append(slice, i)
	}
}

func copyFuncOfSlice() {
	src := []int{6, 3, 2, 4}
	des := make([]int, 10)
	/*返回复制的*/
	quantity := copy(des, src)
	fmt.Println(quantity)
	fmt.Println(des)

	fmt.Println(&src[0])
	fmt.Println(&des[0])
}

func main() {
	capacityOfSlice()
}
