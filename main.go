package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	nums := [10]int{}

	for i, _ := range nums {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		nums[i] = r.Intn(10)
	}

	fmt.Println(nums)

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("the num %v is EVEN \n", num)
		} else {
			fmt.Printf("the num %v is ODD \n", num)
		}
	}
}
