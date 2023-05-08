package main

import "fmt"

func main() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	fmt.Println("----------------------")

	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf(resultString)

	fmt.Println("----------------------")

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("============================")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}
