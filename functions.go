package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusplus(a,b,c int) int {
	return a+b+c
}

func main() {
	result:=plus(1,2)
	fmt.Println("1+2 =", result)

	result2:=plusplus(1,2,3)
	fmt.Println("1+2+3 =", result2)
}