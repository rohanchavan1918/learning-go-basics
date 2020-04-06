package main

import "fmt"

func main() {

	var name string
	fmt.Println("Enter your Name")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	strArr := []string{"Coders", "Arena"}
	for i, c := range strArr {
		fmt.Println(i)
		fmt.Println(c)
	}

	//  Can also range over on map

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("key :", k, " Valve : ", v)
	}
}
