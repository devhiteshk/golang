package main

import "fmt"

func sum(num ...int) {
	// we get slice in the function
	fmt.Println(num, " ")
	total := 0

	for _, v := range num {
		total += v
	}
	fmt.Println(total)
}

func main() {
	sum(1)
	sum(1, 2, 3)
	sum(1, 10)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
