package main 


import "fmt"


func main() {
	// creating a names array of type sting\
	names := [4]string{
	"Anuj",
	"Sam",
	"Dev",
	"Pooper",
	}

	fmt.Println(names)

	// slices dont store any data just describes the section of an underlying array
	// changing the slice element changes the underlying array values 

	a := names[1:3]
	b := names[0:2]

	fmt.Println("this a:", a)
	fmt.Println("This b:",b)

	b[0] = "Changed value"

	fmt.Println("names arrchanged", names)



}
