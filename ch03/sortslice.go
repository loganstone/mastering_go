package main

import (
	"fmt"
	"sort"
)

type person struct {
	name   string
	height int
	weight int
}

func main() {
	people := make([]person, 0)
	fmt.Println(people)

	people = append(people, person{"Logan", 178, 75})
	people = append(people, person{"Super Logan", 188, 90})
	people = append(people, person{"Coma", 162, 45})
	people = append(people, person{"Lovely", 90, 12})
	people = append(people, person{"Jane doe", 172, 55})

	fmt.Println("0:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].height < people[j].height
	})
	fmt.Println("<:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})
	fmt.Println(">:", people)

}
