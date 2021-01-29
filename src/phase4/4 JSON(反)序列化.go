package main

import (
	"encoding/json"
	"fmt"
)

type SaoBi struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Score    float32 `json:"score"`
}

/*
结构体、切片、map都是一个流程
*/
func getJsonStr(i interface{}) string {
	srlBytes, _ := json.Marshal(i)
	srlStr := string(srlBytes)
	return srlStr
}

func structSrl() string {
	saobi := &SaoBi{
		Name:     "林木婷子",
		Age:      18,
		Birthday: "2002-7-4",
		Score:    5.5,
	}
	return getJsonStr(saobi)
}

func mapSrl() string {
	iMap := make(map[string]interface{}, 10)
	saobi := &SaoBi{
		Name:     "shika小鹿鹿",
		Age:      22,
		Birthday: "2003-8-8",
		Score:    6,
	}
	iMap["a"] = saobi
	iMap["b"] = 6324
	iMap["c"] = "骚骚大骚逼"

	return getJsonStr(iMap)
}

func sliceSrl() string {
	slice := make([]string, 0)
	slice = append(slice, "下", "贱", "工", "人")
	return getJsonStr(slice)
}

func unmarshalStruct() {
	srl := structSrl()
	saobi := new(SaoBi)
	err := json.Unmarshal([]byte(srl), saobi)
	if err == nil {
		fmt.Println(*saobi)
	}
}

func unmarshalMap() {
	/*
		反序列化map时，不需要为map指针make空间
	*/
	var m map[string]interface{}
	js := mapSrl()
	err := json.Unmarshal([]byte(js), &m)
	if err == nil {
		fmt.Println(m)
	}
}

func unmarshalSlice() {
	js := sliceSrl()
	var slice []string
	err := json.Unmarshal([]byte(js), &slice)
	if err == nil {
		fmt.Println(slice)
	}
}

func main() {
	unmarshalSlice()
}
