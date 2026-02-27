package calc

import (
	"fmt"
)

var operationsCount int

func init() {
	operationsCount = 0
	fmt.Println("calc package initialized")
}

func Sum(nums ...float64) (result float64) {

	for _, num := range nums {
		result += num
	}
	operationsCount++
	return

}
func Max(nums ...float64) (result float64) {

	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	operationsCount++
	return

}
func Min(nums ...float64) (result float64) {

	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	operationsCount++
	return
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	operationsCount++
	return a / b, nil
}

func GetOperationsCount() int {
	return operationsCount
}

//
// завдання 2
//

type Calc struct {
}

func (c Calc) Sum(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (c Calc) Max(nums ...float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func (c Calc) Min(nums ...float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func (c Calc) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}
