package main

import "fmt"

func main() {
	//算术运算符
	//++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
	//关系运算符
	//逻辑运算符
	//位运算符
	//赋值运算符

	numbers := []int{2, 3, 4, 5, 2, 3, 4}
	singleNumber := findSingleNumber(numbers)
	fmt.Printf("出现一次的数字是：%d\n", singleNumber)

}

func findSingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
