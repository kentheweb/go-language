package main

import (
	"fmt"
)

func main() {
	var maps map[string]int = map[string]int{
		"name": 3,
		"age":  20,
	}
	fmt.Println(maps)
	// adding items in a map
	maps["year"] = 2019

	fmt.Println(maps)
	// deleting keys in an map
	delete(maps, "name")

	// accessing items in a map
	fmt.Println(maps["year"])

	// finding the length of a map
	fmt.Println(`The length of the map is`, len(maps))

	fmt.Println(`the items in a map are :`, maps)
}
