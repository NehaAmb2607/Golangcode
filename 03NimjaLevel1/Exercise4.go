package main

import "fmt"

type sink int

var w sink
var q int

func main() {
	q = 70
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	w = 42
	fmt.Print(w, "\n")
	fmt.Printf("%T\n", w)
	fmt.Print(q, "\n")
	fmt.Printf("%T\n", q)
	q = int(w)
	fmt.Println(q, "\n")
	fmt.Printf("%T\n", q)
}
