package main

import "fmt"

func main() {
	var (
		a int             // using default value
		b string = "hola" // specify the value
		c        = true   // infering the type
	)

	d := 5 // you don't need var in a pure infer notation

	const e = 10 // you don't need type in a const

	fmt.Printf("a is %d and b is %s and c is %t and d is %d", a, b, c, d)
}
