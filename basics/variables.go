package main

import "fmt"

func main() {

	var a int = 5
	fmt.Println(a)

	var b, c int = 6, 7
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e float64 = 3.14
	fmt.Println(e)

	var f int
	fmt.Println(f)

	var g string = "Hello"
	fmt.Println(g)

	h := "World"
	fmt.Println(h)

	var (
		i = 10
		j = 11
		k = "Twelve"
	)
	fmt.Println(i, j, k)

}
