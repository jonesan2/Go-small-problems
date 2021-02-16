package main

import "fmt"

func main() {

	// create an empty map
	m := make(map[string]int)

	// set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// get value
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// return number of key/value pairs
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	// the optional second return value when getting a value is a boolean that
	//     indicates whether the key was present in the map
	// the blank identifier _ is used for ignored return values
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
