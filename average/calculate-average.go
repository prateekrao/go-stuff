package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++

		if n == 0 {
			fmt.Println("No numbers were entered.")
			os.Exit(-1)
		}

	}
	fmt.Println("The average is ", sum/float64(n))
}
