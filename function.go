package main

import "fmt"

func main() {
	var names [3]string = [3]string{"김춘삼", "이순신", "홍길동"}

	for index, name := range names {
		fmt.Println(index, name)
	}
	몫, 성공여부 := Divide(3, 1)
	fmt.Println(몫, 성공여부)
}
func Divide(a, b int) (int, bool) {
	// 젯수가 0이면 나눌 수 없음
	if b == 0 {
		return 0, false
	}
	// 나누기 가능
	return a / b, true
}
