package main

import "fmt"

func main() {
	say("This", "is", "a", "book")
	say("Hi")
}

func say(msg ...string) {
	for index, _ := range msg {
		fmt.Println(index)
	}
}
