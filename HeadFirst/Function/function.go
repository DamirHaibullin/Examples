package main

import "fmt"

func main() {
	paintNeeded(3.2, 5.3)
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed \n", area/10.0)
}
