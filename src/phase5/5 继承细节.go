package main

import "fmt"

/**
1.多重继承：
结构体可以有两个或更多个匿名结构体字段
如果这些结构体字段有同名的字段或方法，但是本结构体又并没有同名的，
那么在使用或访问这些字段或方法时，必须明确指定匿名结构体的名字。

韩老师建议：尽量别用
*/

type SweetGirl struct {
	name      string
	bestDress string
}
type Programmer struct {
	name     string
	bestGame string
}
type MtfCoder struct {
	/*
		2.如果是有名结构体字段，这种模式叫组合，
		当然必须加该字段名才能访问
	*/
	//dreamGirl SweetGirl

	/*
		3.结构体的匿名字段可以是基本数据类型，
		同一基本数据类型匿名字段只能有一个
	*/
	int
	//int
	SweetGirl
	Programmer
}

func main() {
	mCoder := new(MtfCoder)
	//mCoder.name="薇儿" //ambiguous
	mCoder.SweetGirl.name = "薇儿"
	mCoder.Programmer.name = "周深"

	mCoder.bestDress = "粉色蕾丝鱼尾睡裙"
	mCoder.bestGame = "原神"

	//mCoder.dreamGirl.name="方圆"
	//mCoder.dreamGirl.bestDress="黑色包臀"

	//fmt.Println(*mCoder)

	/*
		4.可以在创建时就指定各个匿名结构体字段的值
	*/
	weier := MtfCoder{
		int: 42,
		SweetGirl: SweetGirl{
			name:      "薇儿",
			bestDress: "蓝色碎花吊带公主裙",
		},
		Programmer: Programmer{
			name:     "周瑜",
			bestGame: "原神",
		},
	}
	weier.int = 42
	fmt.Println(weier)
}
