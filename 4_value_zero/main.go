package main

import "fmt"

var name string
var age int
var temperature float32

func main() {

	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", age, age)
	fmt.Printf("%v, %T\n", temperature, temperature)
}
