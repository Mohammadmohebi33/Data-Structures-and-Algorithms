package main

import "fmt"

func digitSum(n int) int {

	if n < 10 {
		return n
	}

	return digitSum(n/10) + (n % 10)
}

func main() {

	// f(167)
	// f(16) + 7
	// f(1) + 6
	// f(1) => 1
	fmt.Println(digitSum(167))
}
