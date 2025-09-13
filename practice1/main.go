package main

import "fmt"

var name string = ""

func main() {
	fmt.Println("Hello, my name is Nurgaliev Dias")

	fmt.Scan(&name)

	fmt.Println("Hello, " + name)
}
