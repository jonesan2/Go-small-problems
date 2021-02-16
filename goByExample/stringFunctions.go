package main

import (
	"fmt"
	// alias the strings package to s for brevity of code
	s "strings"
)

// alias fmt.Println as well
var p = fmt.Println

func main() {
	// a sample of the functions available in the strings package
	// since these are functions from the package and not methods on the type,
	//    we need to pass the string in as the first argument and call the
	//    functions with the package as the receiver
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // negative means no limit
	p("Replace:   ", s.Replace("foo", "o", "0", 1))  // limit 1 replacement
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// These important built-in functions are not part of the strings package
	// These functions work at the byte level. Go uses UTF-8 encoded strings, 
	//    
	p("Len:  ", len("hello"))
	p("Char: ", "hello"[1])
}
