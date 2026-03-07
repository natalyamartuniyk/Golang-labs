package main

import (
	"fmt"
	"sync"
)

func main() {

	filterEven := make(chan int, 2)
	square := make(chan int)
	sumChannel := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {

		defer wg.Done()

		for i := 1; i < 101; i++ {
			filterEven <- i
			//fmt.Println("Numbers: ", i)
		}
		close(filterEven)
	}()

	go func() {

		defer wg.Done()

		for num := range filterEven {
			if num%2 == 0 {
				square <- num
				fmt.Println("Even: ", num)
			}
		}
		close(square)
	}()

	go func() {
		defer wg.Done()

		for num := range square {
			numSquare := num * num
			sumChannel <- numSquare
			fmt.Println("Square: ", numSquare)
		}
		close(sumChannel)
	}()

	go func() {
		defer wg.Done()

		sumOfNumbers := 0
		for num := range sumChannel {
			sumOfNumbers += num
		}
		fmt.Println("Sum of Numbers: ", sumOfNumbers)
	}()

	wg.Wait()

}
