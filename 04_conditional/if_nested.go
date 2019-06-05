package main

import "fmt"

func main() {
	//program to check given no s divisible by 2 as well as 3 or 4
	var no int32 = 20

	if no%2 == 0 {
		if no%3 == 0 {
			fmt.Println("No is diviible by 2 and 3")

		} else if no%4 == 0 {
			fmt.Println("No is diviible by 2 and 4")

		} else {
			fmt.Println("No is diviible by 2 and not divisible by 3 and 4")
		}
	} else {
		fmt.Println("No is odd")
	}

}
