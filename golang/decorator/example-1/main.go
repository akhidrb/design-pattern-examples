package main

import "fmt"

func decorator(f func(string)) func(string) {
	return func(str string) {
		fmt.Println("Before func")
		f(str)
		fmt.Println("after func")
	}
}

func say_whee(str string) {
	fmt.Println(str)
}

func main() {
	decorator(say_whee)("whee")
}