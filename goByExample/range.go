// range iterates over elements in a variety of data structures

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// provides index and value, in that order
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// on a map, range iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// by ignoring the value return value, range can iterate over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// on strings, range iterates over Unicode code points
	// iterates over index, rune(code point)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
