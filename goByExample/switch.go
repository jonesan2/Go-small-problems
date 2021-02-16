package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // commas to separate multiple expressions
		fmt.Println("It's the weekend")
	default: // optional
		fmt.Println("It's a weekday")
	}

	// with no expression, an alternate if/else logic structure
	t := time.Now()
	switch {
	case t.Hour() < 12: // case expressions can be non-constants
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// type switch: compares types instead of values
	// can be used to discover the type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
