package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64 = 30

	result := math.Sqrt(number)

	fmt.Println(math.Ceil(result))
	// using printf for integer manipulation
	fmt.Printf("%.3f\n\n ", result)
	fmt.Printf("%2g", result)

}
