package main 

import "fmt"

func main() {
	// array fixed, slices are dynamically-sized
	// array will always have a fixed lenth like below and slice will be []int{1,2,3}
	primes := [6]int{2,3,5,7,11,13}


	//slices have low and upper bounds seperated by colon a[low:high]
	var slice []int = primes[1:4]
	fmt.Println(slice)

	// slice has both a length and capacity 
	// length of the slice is the elements it contains cap(s)
	// capacity of slice is the number of elements that are in the underlying array len(s)

	s := []int{2,3,5,6,7,13,83} // This is a slice cuz no defined number in the []

	printSlice(s)

	// slicing the slice to give it zero length 

	s = s[:0]
	printSlice(s)

	// extending the slice length 
	s = s[:4] 
	printSlice(s)

	// drop the first 2 values 
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s) , cap(s), s)
}
