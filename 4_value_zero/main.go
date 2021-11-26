package main

import "fmt"

var name string
var age int
var temperature float32
var isNull bool

func main() {

	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", age, age)
	fmt.Printf("%v, %T\n", temperature, temperature)
	fmt.Printf("%v, %T\n", isNull, isNull)
}
