package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(add(23, 32.5))
}

func add(a int, b float64) int {
	return int(a + int(b))
}
