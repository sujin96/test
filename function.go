package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println("호출전")
	fmt.Printf("a = %p, b = %p\n", &a, &b)
	swap(&a, &b)
	fmt.Println("호출후")
	fmt.Printf("a = %d, b = %d\n", a, b)
}
func swap(i *int, j *int) {
	temp := *i
	*i = *j
	*j = temp
	fmt.Printf("i = %d, j = %d\n", *i, *j)
}
