package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func fn1() {
	people := []Person{
		{"Pat", "Patterson", 38},
		{"Tracy", "Tracer", 23},
		{"Anders", "Anderson", 29},
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("Sort by LastName:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("Sort reverse by Age:", people)
}
