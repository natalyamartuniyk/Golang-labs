package main

import "fmt"

type Currency interface {
	Name() string
	USDRate() float64
	IsCrypto() bool
}

type USD struct{}

func (u USD) Name() string {
	return "USD"
}
func (u USD) USDRate() float64 {
	return 1.0
}
func (u USD) IsCrypto() bool {
	return false
}

type Bitcoin struct{}

func (b Bitcoin) Name() string {
	return "Bitcoin"
}
func (b Bitcoin) USDRate() float64 {
	return 64551.65
}
func (b Bitcoin) IsCrypto() bool {
	return false
}

func main() {

	var c1 Currency = USD{}
	var c2 Currency = Bitcoin{}

	fmt.Println(c1.Name(), c1.USDRate(), c1.IsCrypto())
	fmt.Println(c2.Name(), c2.USDRate(), c2.IsCrypto())

	//завдання 1

	/*a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := []int{2, 20, 4, 30, 6, 40, 8, 50, 10, 60}

	result := make([]int, 10)

	for i := 0; i < 10; i++ {

		result[i] = a[i] + b[i]
	}

	fmt.Println(result)

	*/

}
