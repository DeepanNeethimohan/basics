package main

import "fmt"

type myvar int

var x myvar
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y := int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}