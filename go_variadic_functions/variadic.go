package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 4, 5, 56)

	nums := []int{3, 6, 3, 8, 3, 3}
	sum(nums...)
}
