package main

import "fmt"

func stringFunc() {
	var s string = "hello"
	var b byte = s[1] //returns 101
	fmt.Println(b)

	var a string = "learning go"
	var c string = a[2:7]
	fmt.Println(c)
	fmt.Println("Length: ", len(a))

}
