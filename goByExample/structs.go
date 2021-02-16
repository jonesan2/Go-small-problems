package main

import "fmt"

type person struct {
	name string
	age  int
}

// constructor: constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p // you can safely return a pointer to a local variable, as the
	// local variable will survive the scope of the function
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"}) // omitted fields will be zero-valued
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	// you can access struct fields with a dot .
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can use dots with struct pointers: they are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

}
