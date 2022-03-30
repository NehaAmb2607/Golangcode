package main

import "fmt"

const (
	y         = 24
	t float64 = 25.3698
)

func main() {
	a := 12 == 12
	b := 24 <= 25
	c := 24 <= 25
	d := 28 >= 25
	e := 24 < 96
	f := 24 != 0
	g := 40 > 25
	fmt.Println(a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f, "\n", g, "\n", t, "\n", y)

}
