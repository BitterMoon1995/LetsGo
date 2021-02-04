package test

import (
	"testing"
)

func getNigger(i int) int {
	if i > 100 {
		return 8964
	} else if i <= 10 {
		return 6324
	} else {
		return 4396
	}
}

/*
测试例程命名规范：
	包名随意，和被测试函数同包即可。
	go文件名必须为 xxx_test ，即以_test结尾。
	函数名必须为 TestXxx ，即Test开头+被测试函数，后者首字母必须大写。
目录结构：
	测试函数被测试函数同包或同文件，或后者为public，能访问即可
*/
func TestGetNigger(t *testing.T) {
	nigger := getNigger(88)
	if nigger == 8964 {
		/*
			带f的还可以格式化输出
		*/
		t.Fatalf("致命错误")
	} else if nigger == 6324 {
		t.Errorf("错误")
	} else {
		t.Logf("返回值为：%v 测试通过", nigger)
	}
}

func TestGoroutineBreak(t *testing.T) {
	c := make(chan bool, 4)
	go func() {
		for i := 0; i < 4; i++ {
			c <- true
		}
	}()

	for i := 0; i < 4; i++ {
		<-c
	}
	close(c)
	t.Logf("赢了")
}
