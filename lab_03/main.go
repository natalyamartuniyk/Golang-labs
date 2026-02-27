package main

import (
	"fmt"
	"lab_03/calc"
)

func main() {

	sum := calc.Sum(1, 3, 4)
	fmt.Println("Sum: ", sum)

	max := calc.Max(10, 4568, 2, 88888)
	fmt.Println("Max: ", max)

	min := calc.Min(-1, 23490281, -1183026492, 0)
	fmt.Println("Min: ", min)

	dilennya, err := calc.Divide(10, 15)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Divide: ", dilennya)

	fmt.Println("Operations count: ", calc.GetOperationsCount())

	//
	// завдання 2
	//

	var calculator calc.Calculator = calc.Calc{}

	fmt.Println("Sum:", calculator.Sum(1, 2, 3, 4))
	fmt.Println("Max:", calculator.Max(12, 45218, 2, 8320188))
	fmt.Println("Min:", calculator.Min(10, 58211, 310, 7392271))
	result, err := calculator.Divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide: ", result)
	}

}
