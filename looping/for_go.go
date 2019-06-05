package main

import "fmt"

func main() {

	//progrm to print 1 to 10
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//go with a condition
	var b bool = true
	for b {
		fmt.Println(b)
		b = !b
	}
}
