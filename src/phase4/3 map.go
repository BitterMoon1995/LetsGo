package main

import (
	"fmt"
	"sort"
)

/*
map声明后，并不能立即使用，必须先make分配空间
*/
func declareMap() {
	/*仅仅声明了，此时map0=nil*/
	var map0 map[string]string
	map0 = make(map[string]string, 10)
	map0["a"] = "nigga"
	/*相同key，覆盖value*/
	//map0["a"] = "nigga~~~~~"
	map0["b"] = "nigger"
	map0["c"] = "black slave"

	fmt.Println(map0)
}

func fasterInitialize() {
	/*★声明时顺手make了*/
	capitalMap := make(map[string]string, 10)
	capitalMap["USA"] = "New York"
	capitalMap["苏维埃共产主义满洲国"] = "咱老北京儿"
	fmt.Println(capitalMap)
	/*声明时录入初始值*/
	provCapitalMap := map[string]string{
		"四川县": "我大成都",
		"河北村": "八旗老北京儿",
	}
	fmt.Println(provCapitalMap)
}

func crudOfMap() {
	capitalMap := make(map[string]string, 10)
	s := "USA"
	/*增*/
	capitalMap[s] = "New York"
	capitalMap["Hell"] = "北京儿"
	capitalMap["England"] = "London"
	/*改
	如果key已存在，就是修改*/
	capitalMap["Hell"] = "成都"
	/*删
	删不存在的key不会报错。
	删所有元素，要么遍历，要么利用GC*/
	delete(capitalMap, "England")
	delete(capitalMap, "叙利亚")

	/*查
	可接收两个返回值*/
	key := "Hell2"
	value, isExists := capitalMap[key]
	if isExists {
		fmt.Println(value)
	} else {
		fmt.Println("没有key：", key)
	}
	fmt.Println(capitalMap)

	/*长度*/
	fmt.Println("长度：", len(capitalMap))
}

func traverseOfMap() {
	capitals := make(map[string]string, 10)
	capitals["韩国"] = "首尔"
	capitals["日本"] = "东京"
	capitals["美国"] = "华盛顿"

	for i, e := range capitals {
		fmt.Println(i, e)
	}
}

func sliceOfMap() {
	name := "name"
	position := "position"

	girls := make([]map[string]string, 10)

	wei := make(map[string]string, 2)
	wei[name] = "薇儿"
	wei[position] = "美腿校花"
	/*对切片的增加操作必须使用append动态追加*/
	girls = append(girls, wei)

	rui := make(map[string]string, 2)
	rui[name] = "蕊蕊"
	rui[position] = "仙颜校花"
	girls = append(girls, rui)

	fmt.Println(girls)
}

/*
map本身是无序的，go也没有内置的函数来对其排序
可以利用一个缓冲切片以及sort包内的函数对map进行排序
*/
func sortOfMap() {
	intMap := make(map[int]string)
	intMap[10] = "cause"
	intMap[71] = "beautiful"
	intMap[22] = "I`m"
	intMap[88] = "nigger"
	fmt.Println(intMap)

	/*
		用一个切片接收map所有key
		再利用sort.Ints函数对该切片进行排序
		再遍历排序后的切片，作为索引输出map的value
	*/
	buffer := make([]int, 0)
	for i, _ := range intMap {

		buffer = append(buffer, i)
		sort.Ints(buffer)
	}

	fmt.Println(buffer)

	for _, e := range buffer {
		fmt.Println(intMap[e])
	}
}

func main() {
	sortOfMap()
}
