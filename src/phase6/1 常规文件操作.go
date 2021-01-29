package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
流：数据在数据源和程序之间经历的路径
*/
func openFile() {
	file, err := os.Open("F:/文档/memes/麻了2.txt")
	if err != nil {
		fmt.Println("打开失败：", err)
	} else {
		fmt.Println(file.Name())
		closeFile(file)
	}
}

/*
文件指针（描述符）用完后必须关闭，否则
*/
func closeFile(files ...*os.File) {

	for _, file := range files {
		err := file.Close()
		if err != nil {
			fmt.Println(file, "关闭失败：", err)
		}
	}
}

func bufferedRead() {
	file, err := os.Open("F:/文档/memes/麻了.txt")
	if err != nil {
		fmt.Println("打开失败：", err)
	} else {
		reader := bufio.NewReader(file)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Println(str)
		}
	}
	defer closeFile(file)
}

func bufferedWrite(path string) error {
	/*
		windows下FileMode无效
		创建：os.O_RDWR | os.O_CREATE
		重写：os.O_RDWR | O_TRUNC
		追加：os.O_RDWR | os.O_APPEND

		追加和重写一直报Access is denied.
	*/
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("打开出错：", err)
		return err
	}
	str := "骚逼蹄子\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("写入出错：", err)
			return err
		}
		/*
			带缓存的输出流注意写完要flush，让数据从缓存入盘
		*/
		err2 := writer.Flush()
		if err2 != nil {
			fmt.Println("flush出错：", err2)
			return err2
		}
	}
	defer closeFile(file)
	return nil
}

func fileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func fileCopy(src string, dst string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	reader := bufio.NewReader(srcFile)

	/*创建writer前，判断目标文件或目录是否存在，不存在则先创建*/
	exist, err := fileExist(dst)
	var dstFile *os.File
	if exist {
		fmt.Println("文件已存在")
		return 0, nil
	} else {
		dstFile, _ = os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 777)
	}

	writer := bufio.NewWriter(dstFile)

	w, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("copy失败")
		return 0, err
	}
	defer closeFile(srcFile, dstFile)
	return w, nil
}

func main() {
	written, err := fileCopy("F:\\图片\\骚牌\\saobi.jpg", "F:\\图片\\骚牌\\shele\\saobi.jpg")
	fmt.Println(written, err)
}
