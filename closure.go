


// go supports anonymous fn that can be used to form closures
// Closures are functions that remember their outer env variables even if the outer function has finished executing
// function gets "close over" along with its environment (variables, functions that modify the state or var)
//  closure is like a dabba andar saman daldo bahar nikal ke use karo using attached functions to modify the state fir bhi andar ke variables safe rehte hai


package main 

import "fmt"

	// defining func signature here 
	func initSeq() func() int {  // func() int specifies that  initSeq will return a function and the return func will return int
		count := 0

		return func() int {
			count++
			return count
		}
	}

	func main() {
		dabba1 := initSeq()
		dabba2 := initSeq()

		fmt.Println("d1",dabba1())
		fmt.Println("d1", dabba1())

		fmt.Println("d2", dabba2())
		fmt.Println("d2", dabba2())
		
	}
