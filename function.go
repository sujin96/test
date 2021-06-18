package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	var b []int
	//for 사용하여 a찍기
	for i := 0; i < 3; i++ {
		fmt.Print(a[i])
	}
	//range 사용하여 b배열찍기
	for _, name := range b {
		fmt.Print(name)
	}
}
