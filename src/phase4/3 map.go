package main

import "fmt"

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
	/*声明时顺手make了*/
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

func main() {
	fasterInitialize()
}
