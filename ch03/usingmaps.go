package main

import "fmt"

func main() {
	// make
	scores := make(map[string]int)
	scores["reading"] = 50
	scores["writing"] = 70
	scores["math"] = 95
	scores["sciene"] = 99
	scores["essay"] = 100
	fmt.Println("Scores:", scores)

	// literal
	anotherScores := map[string]int{
		"국어":  100,
		"수학":  50,
		"영어":  80,
		"한국사": 55,
		"탐구":  82,
	}
	fmt.Println("Another Scores:", anotherScores)

	delete(anotherScores, "수학")
	delete(anotherScores, "수학") // no error
	delete(anotherScores, "수학") // no error
	delete(anotherScores, "수학") // no error
	delete(anotherScores, "수학") // no error
	fmt.Println("Deleted \"수학\" Another Scores:", anotherScores)

	if v, ok := scores["korean"]; !ok {
		fmt.Println("Does it exist")
	} else {
		fmt.Println("Korean Score is ", v)
	}

	for k, v := range anotherScores {
		fmt.Println(k, v)
	}
}
