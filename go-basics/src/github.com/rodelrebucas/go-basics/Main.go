package main

import "fmt"

// Variables declared at package level should
// always have type
var pkgVar float32 = 1.2

func main() {
	// Declaring variables
	var i int
	i = 42
	var j int = 43
	k := 44 // go infers the type, // used on inside function only

	// Printing
	fmt.Println(i)             // 42
	fmt.Println(j)             // 43
	fmt.Println(k)             // 44
	fmt.Printf("%v, %T", j, j) // 43, int
}
