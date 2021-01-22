package utils

import "fmt"

var Name string
var Age int

func init() {
	fmt.Println("init() of common utils")
	Name = "ts美妖"
	Age = 12
}

/*
盘点十大不需要【public】的理由：
把函数首字母大写表示public不就行了
*/
func GetNigga() {
	fmt.Println(`ah,nigga,don't hate me cause I'm beautiful,nigga.
maybe if you got rid of that old yeeyee ass haircut you got
you'd get some bitches on your dick.
oh better yet,maybe Tanisha'll call your dog ass 
if she ever stop fucking with that brain surgeon or lawyer she fucking with.
nigga~~~~~~~~~~~~~~~~~`)
}

func GeiInfo() {

}
