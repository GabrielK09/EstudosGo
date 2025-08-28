package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	mSlice := array[1:4]

	fmt.Println(mSlice)
	fmt.Println(len(mSlice))
	fmt.Println(cap(mSlice))
}
