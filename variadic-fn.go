

package main

import "fmt"

// variadic functions are functions that can be called with any number of trailing arguments 
// every nums ...int meaning every number will come into the func as a slice []int
func sum(nums ...int) {

	fmt.Print(nums, " ") 
  total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)


}


func main() {
	
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...) // passing the slice of numbs using ...

}




