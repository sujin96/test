package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(a, b)
}
func swap(i int, j int) {
	temp := i
	i = j
	j = temp
	fmt.Printf("i = %d, j = %d\n", i, j)
}
