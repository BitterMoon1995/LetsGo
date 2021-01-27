package main

import (
	"fmt"
)

type USB interface {
	/*
		接口中只能有未实现的方法，不能有任何变量
	*/
	//Name string
	insert()
	pullout()
	/*
		一个结构体实现了一个接口所有方法才叫实现了这个接口，
		如果只实现了部分方法，该结构体不具备该接口的多态特性
	*/
	//nigger()
}
type Phone struct {
	Name string
}

func (p Phone) insert() {
	fmt.Println("手机插入USB")
}

func (p Phone) pullout() {
	fmt.Println("手机拔出USB")
}

func (p Phone) shootForTikTok() {
	fmt.Println("拍发骚视频发抖音")
}

type Camera struct {
	Name string
}

func (c Camera) insert() {
	fmt.Println("相机插入USB")
}

func (c Camera) pullout() {
	fmt.Println("相机拔出USB")
}

type Computer struct {
}

func (computer Computer) supportUSB(usb USB) {
	fmt.Println("【本机支持USB接口】")
	usb.insert()
	usb.pullout()
}

func main() {
	phone := new(Phone)
	camera := new(Camera)

	computer := new(Computer)
	computer.supportUSB(phone)
	computer.supportUSB(camera)
}
