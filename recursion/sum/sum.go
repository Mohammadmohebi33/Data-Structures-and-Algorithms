package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return sum(n-1) + n
}

func main() {
	fmt.Println(sum(5))
}
