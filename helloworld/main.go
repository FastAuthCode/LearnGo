package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("hello world")
	randmo := rand.Intn(9) + 1
	fmt.Println(randmo)
	time.Sleep(time.Duration(randmo) * time.Second)
	end := time.Now()
	var cost = end.Sub(start)
	fmt.Println(cost)

	fmt.Println("Days of the week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

}
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		s := target - num

		if index, ok := indexMap[s]; ok {
			return []int{index, i}
		}

		indexMap[num] = i
	}

	return []int{}
}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
