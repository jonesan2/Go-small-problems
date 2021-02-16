package main

import (
	"fmt"
	"os"
	"reflect"
)

type Saidan struct {
	Name  string
	Power int
}

func (s *Saidan) Super() {
	s.Power += 10000
}

func main() {

	gok := Saidan{"Sasuke", 6999}
	goku := &gok
	// goku := &Saidan{"Naruto", 7000}
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[0])
	fmt.Println(reflect.TypeOf(goku))
	fmt.Println(reflect.TypeOf(*goku))
	fmt.Println(reflect.TypeOf(&goku))

	goku.Super()
	fmt.Println(*goku)
	fmt.Println(gok)
	goku.Power += 10000
	fmt.Println(*goku)
	fmt.Println(gok)
	gok.Power += 10000
	fmt.Println(*goku)
	fmt.Println(gok)
	(&gok).Power += 10000
	fmt.Println(*goku)
	fmt.Println(gok)
}
