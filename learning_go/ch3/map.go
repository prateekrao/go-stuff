package main

import "fmt"

func mapsFunc() {
	var nilMap map[string]int
	fmt.Println("nilmap:", nilMap)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Zoe"},
	}
	fmt.Println("Teams:", teams)

	ages := make(map[int][]string, 10)
	fmt.Println(ages)

	//Reading and Writing a Map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins)
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Lions"])
	fmt.Println(totalWins["Kittens"])
	fmt.Println("Kittens:", totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println("Kittens:", totalWins["Kittens"])
}
