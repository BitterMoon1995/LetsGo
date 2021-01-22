package main

import (
	"fmt"
	"strings"
)

func multiLine() {
	s := `black
nigger
bees`
	fmt.Println(s)
}

func usualMethods() {
	var str = `ah,nigga,don't hate me cause I'm beautiful,nigga.
maybe if you got rid of that old yeeyee ass haircut you got
you'd get some bitches on your dick.
oh better yet,maybe Tanisha'll call your dog ass 
if she ever stop fucking with that brain surgeon or lawyer she fucking with.
nigga~~~~~`

	//fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println("ah" + "nigga")
	split := strings.Split(str, "nigga")
	sprint := fmt.Sprintln("get", "some", "bitches", "on", "your", "dick")
	fmt.Println(sprint)

	fmt.Println(split)
	fmt.Println(strings.Contains(str, "nigga"))
	fmt.Println(strings.HasSuffix("fuck rainy.avi", "avi"))

	fmt.Println(strings.Index(str, "nigga"))
	fmt.Println(strings.LastIndex(str, "nigga"))
	fmt.Println(strings.Index("fairy", "p"))

	var strList = []string{"don`t", "hate", "me", "cause", "I`m beautiful", "nigga"}
	join := strings.Join(strList, "-")
	fmt.Println(join)
}

func main() {
	usualMethods()
}
