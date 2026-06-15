package s1

import "fmt"

func printNumber(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	printNumber(n - 1)
	fmt.Println(n)
}

//
//func main() {
//	printNumber(5)
//}
