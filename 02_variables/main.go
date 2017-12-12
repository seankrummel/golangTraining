package main

import "fmt"

func main() {

	// shorthand declare and assign
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t %T \n", a, a)
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
