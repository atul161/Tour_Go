package main

import "fmt"

func main() {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr[0])
	fmt.Println(arr[1])

	///dynamic
	a := [2]int{12, 34}
	fmt.Println(a[0])
	fmt.Println(a[1])
	//slicing
	slicer := []string{"hel", "oat", "peak"}
	fmt.Println(slicer[1:2])

}
