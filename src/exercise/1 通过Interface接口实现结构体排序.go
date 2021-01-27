package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Hero struct {
	Name   string
	Damage int
}

/*
基于一个已有的类型，将其作为新类型的类型说明
*/
type HeroSlice []Hero

type Integer int

func (h HeroSlice) Len() int {
	return len(h)
}

func (h HeroSlice) Less(i, j int) bool {
	return h[i].Damage < h[j].Damage
}

func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func sortTheHeroes() {
	heroes := HeroSlice{}
	klee := Hero{
		Name:   "可莉小宝宝",
		Damage: 10,
	}
	diluc := Hero{
		Name:   "老卢",
		Damage: 8,
	}
	ganyu := Hero{
		Name:   "甘雨",
		Damage: 9,
	}
	heroes = append(heroes, klee, diluc, ganyu)
	fmt.Println(heroes)

	sort.Sort(heroes)
	fmt.Println(heroes)
}

func main() {
	var i Integer = 10
	i = 88
	fmt.Println(reflect.TypeOf(i))
}
