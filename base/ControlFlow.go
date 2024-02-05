package main

import "fmt"

func main() {

	//goto语句通过标签进行代码间的无条件跳转。

	gotoDemo1()
	gotoDemo2()
	MultiplicationTable()
}

func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			fmt.Println("结束for循环")
			break
		}
	}
}

func gotoDemo2() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakFlag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		return
	}
breakFlag:
	fmt.Println("结束for循环")
}

func MultiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", i, j, i*j)
		}
		fmt.Println()

	}
}
