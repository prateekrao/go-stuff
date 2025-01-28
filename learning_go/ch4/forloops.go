package main

import "fmt"

func forloops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	for a := 0; a < 10; {
		fmt.Println(a)
		if a%2 == 0 {
			a++
		} else {
			a += 2
		}
	}

	//for-range loops
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	maps := map[int]string{
		1: "hello",
		2: "world",
		3: "golang",
	}
	for k, v := range maps {
		fmt.Println(k, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	//strings in for range loops
	samples := []string{"hello", "apple"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

}
