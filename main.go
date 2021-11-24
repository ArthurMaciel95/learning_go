package main

import "fmt"

func main() {

	/* showMessage("bom dia") */

	sumNumbers(2, 4)

}

func showMessage(message string) {
	fmt.Println(message)
}

func sumNumbers(x int, y int) {
	sum := x + y

	fmt.Println(sum)
}
