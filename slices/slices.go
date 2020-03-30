package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println("s initialized as , ", s)
	s[1] = 2
	s[2] = 4
	sLength := len(s)
	fmt.Println(sLength)
	for j := 0; j <= sLength-1; j++ {
		fmt.Println(j)
		s[j] = j
		fmt.Println(s)
	}
	fmt.Println("New S is ", s)
	fmt.Println(s[1:8])
}
