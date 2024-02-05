package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	Print()
	FPrint()
	Sprint()
	Errorf()
	Placeholder()
	Scan()
	Scanf()

}

func Scanf() {
	var (
		name    string
		age     int
		married bool
	)
	_, err := fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)

	if err != nil {
		fmt.Printf("扫描结果失败%s", err)
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

}

func Scan() {
	var (
		name    string
		age     int
		married bool
	)
	_, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		fmt.Printf("扫描结果失败%s", err)
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

}

func Placeholder() {

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)

	i := 12.34
	fmt.Printf("%f\n", i)
	fmt.Printf("%9f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%9.2f\n", i)
	fmt.Printf("%9.f\n", i)

	s1 := "小王子"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%5s\n", s1)
	fmt.Printf("%-5s\n", s1)
	fmt.Printf("%5.7s\n", s1)
	fmt.Printf("%-5.7s\n", s1)
	fmt.Printf("%5.2s\n", s1)
	fmt.Printf("%05s\n", s1)

}

func Errorf() {
	err := fmt.Errorf("这是一个错误")
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)

	fmt.Println(err)
	fmt.Println(w)
	fmt.Println(e)

}

func Print() {
	fmt.Print("在终端打印该信息。")
	name := "沙河小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

func FPrint() {
	filePath := "example.txt"
	_, err := fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	if err != nil {
		return
	}
	fileObj, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件句柄中写入内容
	_, err = fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
	if err != nil {
		return
	}
	fmt.Println("开始读取文件")
	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}
	// 延迟执行关闭文件操作
	err = fileObj.Close()
	if err != nil {
		fmt.Println("关闭文件时发生错误:", err)
		return
	}
	fmt.Printf("文件内容:\n%s", file)

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("删除文件时发生错误:", err)
		return
	}

	fmt.Println("文件已经删除")

}

func Sprint() {
	fmt.Println()
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

}
