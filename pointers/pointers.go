package main

import "fmt"

func main() {
	val := 20
	fmt.Println(val)

	// Address is refered using the &
	ptr := &val
	fmt.Println(ptr)

	// print the values stored at that address
	// * is the defrefrencing operator
	fmt.Println(*ptr)

	*ptr = 10 // lets update the value , using pointer
	fmt.Println("After updating val :-", val)

	val = 50 // update the value of val
	fmt.Println("After updating pointer :-", val)
}
