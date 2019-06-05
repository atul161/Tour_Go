package main

import "fmt"

func main() {
	//program to check given no is even or odd
	var a = []int64{1, 2, 3, 4}
	var sum int64
	for _, value := range a {
		sum += value
	}
	fmt.Println(sum)
}
