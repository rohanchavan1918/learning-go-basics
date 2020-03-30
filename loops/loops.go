package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
		for n := 3; n >= 0; n-- {
			fmt.Println(n)
		}
	}
}
