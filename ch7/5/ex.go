package main

import "fmt"

func Divide(a, b int) (result int, success bool) { // 변수명 정해두고 쓰기
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
}
