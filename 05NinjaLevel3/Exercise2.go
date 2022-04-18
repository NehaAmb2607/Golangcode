package main

import "fmt"

func main() {
	for i := 65; i < 67; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
	    }
    }

}
