package main

import (
	"fmt"
	"strconv"
)

/**
类型转换尽量小转大
*/

/*
建议整型和浮点型运算时，将整型转换为浮点型，而非相反，否则会丢失精度
*/
func intToFloat() {
	var i int = 32
	var f float32 = 4.89
	fmt.Println(float32(i) + f)
}

func othersToString() {
	var (
		i  int     = 68
		f  float64 = 12.356
		b  bool    = true
		by byte    = 'b'
		r  rune    = '马'
	)
	formatInt := strconv.FormatInt(int64(i), 10)
	formatFloat := strconv.FormatFloat(f, 'f', 3, 64)
	formatBool := strconv.FormatBool(b)
	formatUint := strconv.FormatUint(uint64(by), 10)
	formatRune := strconv.FormatUint(uint64(r), 16)

	fmt.Println(formatInt, formatFloat, formatBool, formatUint, formatRune)
}

func stringToOthers() {
	var (
		intStr   string = "4397"
		floatStr string = "2800.4396"
	)
	parseInt, err := strconv.ParseInt(intStr, 10, 64)
	fmt.Printf("%v---%T", parseInt, parseInt)
	fmt.Println(err)

	parseFloat, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println(parseFloat)

}

func main() {
	stringToOthers()
}
