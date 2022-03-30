package main

import "fmt"

func main() {
	a := 18
	fmt.Printf("%d\t%b\t%0x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%0X", b, b, b)
}
