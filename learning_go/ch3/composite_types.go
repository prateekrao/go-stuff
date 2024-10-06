package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Arrays:")
	arrays()
	fmt.Println("Slices:")
	slice()
	fmt.Println("Slicing:")
	slicing()
	fmt.Println("Strings:")
	stringFunc()
	fmt.Println("Maps:")
	mapsFunc()
}

func arrays() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a[0], a[1], a[2])
	var b = [...]int{10, 20, 30}
	fmt.Println(b[0], b[1], b[2])
	fmt.Println(a == b)
}

func slice() {
	var x = []int{10, 20, 30}
	fmt.Println(x[0], x[1], x[2])

	var y []int
	fmt.Println(y == nil)

	var a = []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(a, b))
	var c = []string{"a", "b", "c"}
	d := []string{"d", "e"}
	fmt.Println(len(c))
	a = append(a, 5, 6)
	fmt.Println(a)
	c = append(c, d...)
	fmt.Println(c)
	fmt.Println(cap(c))

	p := make([]string, 5)
	fmt.Println(p)

	p[1] = "Hello"
	p[2] = "World"
	p[3] = "Golang"

	fmt.Println(p)
	clear(p)
	fmt.Println(p)
}

func slicing() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
