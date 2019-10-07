package main

import (
	"fmt"
)

// type vertex struct {
// 	Lat, Long float64
// }

func manin() {
	type vertex int
	var m = map[string]vertex{"there" : 6}
	// m = make(map[string]vertex)
	// m["Bell Labs"] = vertex{
	// 	40.68433, -74.39967,
	// }
	m["here"] = 5
	fmt.Println(m, m["there"])

	fmt.Println(wordCount("Nathnael"))
	
}

type value string
func wordCount(word string) map[string]value {
	
	var wordMap map[string]value
	wordMap = make(map[string]value)
	for i, v := range word {
		// fmt.Println(string(v), i)
		var key string = string(i)
		wordMap[key] = "letter"
		fmt.Printf("%T \n", string(v))
		// fmt.Printf("%T \n", v)
	}
	return wordMap
}