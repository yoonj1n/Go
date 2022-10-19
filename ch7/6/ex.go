package main

import "fmt"

func printNo(a int) {
	if a == 0 {
		return
	}
	fmt.Println(a)
	printNo(a - 1)
	fmt.Println("After", a)
}

func main() {
	printNo(3)
}
