package main

import "fmt"

var x = 100 // package level scope

func main() {

	f1()
	// function expression
	f2 := func() /* anonymous function */ {
		fmt.Printf("%v \n", x)
	}
	f2()

	// shorthand declare and assign, block level scope
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t %T \n", a, a) // %v is value, %T is type
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, d)

	// zero value
	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
}

func f1() {
	fmt.Printf("%v \n", x)
}
