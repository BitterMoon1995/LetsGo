package main

import "fmt"

/**
go不鼓励封装，算了

继承的写法是加入匿名结构体字段

即使父类的属性是小写的(AKA私有的)，子类依然能够继承；
如果不是匿名的结构体，是有名的，就不是继承，叫组合

*/
type Student struct {
	Name  string
	Age   int
	Score float32
}

type Pupil struct {
	Student
}

type Undergraduate struct {
	Age int
	Student
}

func (p *Student) showInfo() {
	fmt.Println(*p)
}

func (p *Student) setScore(score float32) {
	p.Score = score
}

func (p *Student) examining() {
	fmt.Println("【未知年龄段学生】", p.Name, " 考试中，请勿打扰")
}

func (p *Pupil) examining() {
	fmt.Println("【小学生】", p.Name, " 考试中，请勿打扰")
}

func (u *Undergraduate) examining() {
	fmt.Println("【大学生】", u.Name, " 考试中，请勿打扰")
}

func main() {
	animal := new(Pupil)
	animal.Name = "永兴小畜生雄雄"
	animal.Age = 8

	animal.setScore(0.25)
	animal.showInfo()
	animal.examining()

	five := &Undergraduate{}
	five.Name = "成信院废物小杨"
	five.Age = 21

	five.setScore(59)
	five.showInfo()
	five.examining()
	/*
		如果存在与匿名结构体同名的字段或方法
		可以通过.匿名结构体名来区分，否则编译器采取（继承链）就近原则进行访问
	*/
	five.Student.Age = 9000
	fmt.Println(*five)
	five.Student.examining()
}
