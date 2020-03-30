package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["name"] = "Rohan"
	map1["bio"] = "Programming is cool"
	map1["address"] = "Somewhere in Heaven"

	stringKey := make(map[string]int)
	stringKey["name"] = 1
	stringKey["bio"] = 2
	stringKey["address"] = 3

	intKey := make(map[int]string)
	intKey[1] = "rohan"
	intKey[2] = "akshay"
	intKey[3] = "coders arena"

	for i := 0; i <= 10; i++ {
		fmt.Println(map1)
		fmt.Println(stringKey)
		fmt.Println(intKey)
		fmt.Println("----------------")
	}

}
