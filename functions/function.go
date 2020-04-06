package main

import (
	"fmt"
)

//  func <name>(var_name1 type, var_name2 type) return_type {}
func add(a int, b int) int {
	return a + b
}

func profile(name string) {
	pro := "Your Name is " + name + " Your age is " + name
	fmt.Println(pro)
	// return pro
}

// Variadic functions ... Pass any number of inputs to function
func multi(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

// Multi return functions
func calculator(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func main() {
	var a, b int
	// var name string
	fmt.Print("Enter A ")
	fmt.Scan(&a)
	fmt.Print("Enter B ")
	fmt.Scan(&b)

	fmt.Println("[+] Calculator Function")
	add, sub, mult, div := calculator(a, b)
	fmt.Println("[+] Addition is :- ", add)
	fmt.Println("[+] Substraction is :- ", sub)
	fmt.Println("[+] Multiplication is :- ", mult)
	fmt.Println("[+] Division is :- ", div)

	fmt.Println("[+] Variadic function")
	fmt.Println(multi(3, 5, 7, 3, 2))
}
