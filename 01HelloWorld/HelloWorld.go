package main

import "fmt"

func main() {
	fmt.Println("Good morning")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Welcome to foo")
}
