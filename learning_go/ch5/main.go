package main

import (
	"fmt"
	"os"
)

func main() {
	result, remainder, err := divide(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Division Result- res, rem:", result, remainder)

	fmt.Println("Sort a Slice:")
	fn1()
}
