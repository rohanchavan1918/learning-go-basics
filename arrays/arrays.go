package main

import "fmt"

func main() {
	fmt.Println("1D array")
	var intArr [4]int
	fmt.Println(intArr)

	var doubleArray [5][5]int
	doubleArray[0][1] = 5
	doubleArray[0][2] = 6
	fmt.Println(doubleArray)
}
