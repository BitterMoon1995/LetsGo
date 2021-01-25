package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Square struct {
	leftTop    Point
	lowerRight Point
}

type Rectangle struct {
	leftTop    *Point
	lowerRight *Point
}

/*
结构体的字段在内存中是连续分布的
如果字段是指针，那么这些指针字段本身的字段仍是连续的，但是它们所指向的地址就不一定再是连续的了
*/
func sequentialRamDistributed() {
	square := Square{
		leftTop:    Point{1, 3},
		lowerRight: Point{3, 1},
	}
	fmt.Printf("正方形：%p %p %p %p\n", &square.leftTop.x, &square.leftTop.y,
		&square.lowerRight.x, &square.lowerRight.y)

	rectangle := Rectangle{
		leftTop: &Point{
			x: 1,
			y: 3,
		},
		lowerRight: &Point{
			x: 5,
			y: 1,
		},
	}
	fmt.Printf("长方形：%p %p %p %p\n", &rectangle.leftTop.x, &rectangle.leftTop.y,
		&rectangle.lowerRight.x, &rectangle.lowerRight.y)
	fmt.Printf("长方形两点地址：%p %p ", rectangle.leftTop,
		rectangle.lowerRight)
}

/*
结构体是用户单独定义的数据类型，如果两个结构体需要相互转换，
字段的个数、名称、类型必须完全相同
*/
type Position struct {
	x int
	y int
	//z int
}

func structTypeTransform() {
	point := Point{
		x: 1,
		y: 2,
	}
	position := Position(point)
	fmt.Println(position)
}

/*
可以利用type关键字对结构体进行重新定义，相当于重命名
会被考虑为新的数据类型，但可以和原结构体相互转换
*/
type MyPoint Point

func structReDefine() {
	myPoint := MyPoint{
		x: 2,
		y: 2,
	}
	var point Point
	point = Point(myPoint)
	fmt.Println(point)
}

/*
结构体字段一般都得首字母大写，不然其他包访问不到，
这样转成JSON过后前端不太舒服
使用json包 + json tag(结构体标签)对结构体变量进行JSON序列化
底层是反射
*/
type Kitty struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func structToLowercaseJson() {
	cat := Kitty{
		Name:  "小坏喵喵",
		Breed: "美短",
		Age:   1,
	}

	jsonBytesCat, _ := json.Marshal(cat)
	jsonStrCat := string(jsonBytesCat)
	fmt.Println(jsonStrCat)

}

func main() {
	structToLowercaseJson()
}
