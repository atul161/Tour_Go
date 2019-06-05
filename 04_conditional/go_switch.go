package main

import "fmt"

func main() {
	//program to check given no s divisible by 2 as well as 3 or 4
	var no int32 = 20
	var ans string = "not found"

	switch {
	case no == 1:
		ans = "one"
	case no == 2:
		ans = "two"
	case no == 3:
		ans = "three"
	case no == 4:
		ans = "four"
	case no == 5:
		ans = "five"
	default:
		ans = "may be negative or greater then 5"

	}
	fmt.Printf("%q", ans)

}
