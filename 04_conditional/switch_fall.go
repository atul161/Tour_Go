package main

import "fmt"

func main() {
	//program to check given no s divisible by 2 as well as 3 or 4
	no := 4

	var ans string = "not found"

	switch no {
	case 1:
		ans = "one"
	case 2:
		ans = "two"
	case 3:
		ans = "three"
	case 4:
		ans = "four"
	case 5:
		ans = "five"
	default:
		ans = "may be negative or greater then 5"

	}
	fmt.Printf("%q", ans)

}
