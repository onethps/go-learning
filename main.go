package main

import (
	"fmt"
	"math"
)

func main() {
	var height = 1.75
	var weight float64 = 75
	var bmi = weight / math.Pow(height, 2)
	fmt.Print(bmi)
}
