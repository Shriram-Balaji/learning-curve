package main

// format package
import (
	"fmt"
)

var a = "Hello World"

// integer
var c = 20

// boolean
var d = true

// variable declared with a type, are zero-valued
var e int
var f string

func main() {
	// constants are declared with const (duh!)
	const pop = 500000000

	// for is the  only looping construct available in GoLang
	for i := 1; i < pop; i++ {
		if i%4 == 0 {
			output := fmt.Sprintf("%s %d", "Breaking at", i)
			fmt.Println(output)
			break
		}
	}

	// shorthand for init and declaration
	fmt.Println(pop)
}
