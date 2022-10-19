package main

import "fmt"

func Divide(a, b int) (int, bool) { // first () == 매개변수 second () == 리턴
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
}
