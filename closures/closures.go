package main

// closures are anonymous functions
import "fmt"

func return_anon_func() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {
	// Anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous") // PAss the values from here .. in parenthesis

	print_func := return_anon_func()
	print_func("Hello from Anonymous function")
}
