package main

import "fmt"

func reverse(s string) string {
	if len(s) == 1 {
		return s
	}

	return reverse(s[1:]) + string(s[0])
}

func main() {
	fmt.Println(reverse("hello"))
}
